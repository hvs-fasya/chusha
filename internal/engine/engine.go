package engine

import "github.com/hvs-fasya/chusha/internal/models"

// DBInterface - stores common interface
type DBInterface interface {
	TabsGet(bool) ([]*models.Tab, error)
}
