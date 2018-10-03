package engine

import (
	"strconv"

	"github.com/lib/pq"

	"github.com/hvs-fasya/chusha/internal/models"
)

//TabsGet get tabs list
func (db *PgDB) TabsGet(enabled bool) ([]*models.Tab, error) {
	var tabs []*models.Tab
	query := `SELECT t.id, t.title, t.user_type_visible, t.enabled,
				tp.id, tp.type
			FROM tabs t 
			JOIN tab_types tp ON tp.id=t.tab_type_id
			WHERE t.enabled = $1::boolean`
	rows, err := db.Conn.Query(query, strconv.FormatBool(enabled))
	if err != nil {
		return tabs, err
	}
	defer rows.Close()
	for rows.Next() {
		tab := new(models.Tab)
		tab.TabType = new(models.TabType)
		if err := rows.Scan(
			&tab.ID,
			&tab.Title,
			pq.Array(&tab.UserTypeVisible),
			&tab.Enabled,
			&tab.TabType.ID,
			&tab.TabType.Type,
		); err != nil {
			return tabs, err
		}
		tabs = append(tabs, tab)
	}
	return tabs, err
}
