package models

/*Relation modelo para grabar la relación de un usuario con otro */
type Relation struct {
	UserID          string `bson:"userId" json:"userId"`
	UserFollowingID string `bson:"UserFollowingID" json:"UserFollowingID"`
}
