package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserInfo struct {
	Id            *primitive.ObjectID `bson:"_id,omitempty" json:"Id,omitempty"`
	UserId        *primitive.ObjectID  `bson:"user_id,omitempty" json:"user_id,omitempty"`
	FirstName     string              `bson:"first_name" json:"first_name,omitempty" binding:"required"`
	LastName      string              `bson:"last_name,omitempty" json:"last_name,omitempty" binding:"required"`
	UserName      string              `bson:"username,omitempty" json:"username,omitempty"`
	Number        *int64              `bson:"number,omitempty" json:"number,omitempty" binding:"required"`
	Mail          string              `bson:"mail,omitempty" json:"mail,omitempty" binding:"required"`
	IsDeleted     bool                `bson:"is_deleted,omitempty" json:"is_deleted,omitempty"`
	CreatedAt     string              `bson:"created_at,omitempty" json:"created_at,omitempty"`
	UpdatedAt     string              `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
	Age           string              `bson:"age,omitempty" json:"age,omitempty" binding:"required"`
	DeviceDetails *DeviceDetails      `bson:"device_details,omitempty" json:"device_details,omitempty"`
	Lat           *float64            `bson:"lat,omitempty" json:"lat,omitempty" binding:"required"`
	Lng           *float64            `bson:"lng,omitempty" json:"lng,omitempty" binding:"required"`
	Gender        string              `bson:"gender,omitempty" json:"gender,omitempty" binding:"required"`
	Education     string              `bson:"education,omitempty" json:"education,omitempty"`
	Password      string              `bson:"password,omitempty" json:"password,omitempty"`
	AccessToken  string               `bson:"access_token,omitempty" json:"access_token,omitempty"`
}

type DeviceDetails struct {
	BatLvl        *int8               `bson:"bat_lvl,omitempty"           json:"bat_lvl,omitempty"`
	DeviceOS      string              `bson:"device_os,omitempty"         json:"device_os,omitempty"`
	Modalname     string              `bson:"modalname,omitempty"         json:"modalname,omitempty"`
	AppVersion    string              `bson:"app_version,omitempty"       json:"app_version,omitempty"`
	DeviceType    string              `bson:"device_type,omitempty"       json:"device_type,omitempty"`
	StoreVersion  string              `bson:"store_version,omitempty"     json:"store_version,omitempty"`
	StoreBundleID string              `bson:"store_bundle_id,omitempty"   json:"store_bundle_id,omitempty"`
	UserId        *primitive.ObjectID `bson:"user_id,omitempty"           json:"user_id,omitempty"`
}

type LoginCreds struct {
	Username string `bson:"username,omitempty" json:"username"`
	Password string `bson:"password,omitempty" json:"password"`
}


