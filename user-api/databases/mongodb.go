/*
 * @File: databases.mongodb.go
 * @Description: Handles MongoDB connections
 * @Author: Hua Van Thong (huavanthong14@gmail.com)
 */
package databases

import (
	"time"

	"../common"
	"../models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// MongoDB manages MongoDB connection
type MongoDB struct {
	MgDbSession  *mgo.Session
	Databasename string
}

// Init initalizes mongo database
func (db *MongoDB) Init() error {
	db.Databasename = common.Config.MgDbName

	// DialInfo holds options for establishing a session with a MongoDB cluster.
	dialInfo := &mgo.DialInfo{
		Addrs:    []string{common.Config.MgAddrs}, // Get HOST + PORT
		Timeout:  60 * time.Second,
		Database: db.Databasename,            // Database name
		Username: common.Config.MgDbUsername, // Username
		Password: common.Config.MgDbPassword, // Password
	}

	// Create a session which maintains a pool of socket connections
	// to the DB MongoDB database.
	var err error
	db.MgDbSession, err = mgo.DialWithInfo(dialInfo)

	if err != nil {
		panic("Can't connect to mongo, go error: ")
		return err
	}

	return db.initData()
}

// InitData initializes default data
func (db *MongoDB) initData() error {
	var err error
	var count int

	// Check if user collection has at least one document
	sessionCopy := db.MgDbSession.Copy()
	defer sessionCopy.Close()

	// Get a collection to execute the query against.
	collection := sessionCopy.DB(db.Databasename).C(common.ColUsers)
	count, err = collection.Find(bson.M{}).Count()

	if count < 1 {
		// Create admin/admin account
		var user models.User
		user = models.User{bson.NewObjectId(), "admin", "admin"}
		err = collection.Insert(&user)
	}

	return err
}

// Close the existing connection
func (db *MongoDB) Close() {
	if db.MgDbSession != nil {
		db.MgDbSession.Close()
	}
}