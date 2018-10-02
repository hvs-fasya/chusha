package engine

import "github.com/hvs-fasya/chusha/internal/models"

// DBInterface - stores common interface
type DBInterface interface {
	TabsGet() ([]*models.Tab, error)
}
