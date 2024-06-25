package dao

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UserDataAG(username string) primitive.A {
	return bson.A{
		bson.D{{Key: "$match", Value: bson.D{{Key: "username", Value: username}}}},
		bson.D{
			{Key: "$lookup",
				Value: bson.D{
					{Key: "from", Value: "additional_details"},
					{Key: "localField", Value: "user_id"},
					{Key: "foreignField", Value: "user_id"},
					{Key: "as", Value: "device_details"},
				},
			},
		},
		bson.D{{Key: "$unwind", Value: "$device_details"}},
		bson.D{
			{Key: "$project",
				Value: bson.D{
					// {Key: "password", Value: 0},
					{Key: "_id", Value: 0},
					{Key: "created_at", Value: 0},
					{Key: "updated_at", Value: 0},
					{Key: "device_details._id", Value: 0},
					{Key: "device_details.user_id", Value: 0},
				},
			},
		},
	}
}

func UserDataByUserId(userID string) primitive.A {

	var pipeline bson.A
    if userID != "" {
		objectID, err := primitive.ObjectIDFromHex(userID)
		if err != nil {
			panic(err)
		}
		// Add the $match stage to the pipeline
		matchStage := bson.D{{Key: "$match", Value: bson.D{{Key: "user_id", Value: objectID}}}}
		pipeline = append(pipeline, matchStage)
	}
	pipeline = append(pipeline, bson.D{
		{Key: "$lookup",
			Value: bson.D{
				{Key: "from", Value: "additional_details"},
				{Key: "localField", Value: "user_id"},
				{Key: "foreignField", Value: "user_id"},
				{Key: "as", Value: "device_details"},
			},
		},
	},
		bson.D{{Key: "$unwind", Value: "$device_details"}},
		bson.D{
			{Key: "$project",
				Value: bson.D{
					{Key: "_id", Value: 0},
					{Key: "created_at", Value: 0},
					{Key: "updated_at", Value: 0},
					{Key: "device_details._id", Value: 0},
					{Key: "device_details.user_id", Value: 0},
					{Key: "password", Value: 0},
				},
			},
		})
	return pipeline

}

func UpdateUser(accessToken string) (bson.D) {
	return bson.D{{Key: "$set", Value: bson.D{{Key: "access_token", Value: accessToken}}}}
}
