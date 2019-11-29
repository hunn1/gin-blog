package migrate

import (
	"Kronos/app/models"
)

var Models = []interface{}{
	&models.Article{}, &models.User{},
}
