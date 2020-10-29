/*
 * @Description: https://github.com/crazybber
 * @Author: Edward
 * @Date: 2020-09-16 00:47:40
 * @Last Modified by: Eamon
 * @Last Modified time: 2020-09-16 00:47:40
 */

package db

import (
	"github.com/micro-community/micro-starter/models"
)

func migrate() {

	//Migrate the schema
	db.AutoMigrate(&models.User{})

	// Create
	db.Create(&models.User{Key: "D42", ID: 100})

	// Read
	var user models.User

	db.First(&user, 1)                 // find product with integer primary key
	db.First(&user, "code = ?", "D42") // find product with code D42

	// Update - update product's price to 200
	db.Model(&user).Update("Price", 200)
	// Update - update multiple fields
	db.Model(&user).Updates(models.User{ID: 200, Key: "F42"}) // non-zero fields
	db.Model(&user).Updates(map[string]interface{}{"ID": 200, "Key": "F42"})
	// Delete - delete product
	db.Delete(&user, 1)

}
