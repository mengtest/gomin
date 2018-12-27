package model

type DeleteUserReq struct {
	UserIds []int
}

type DisableUserReq struct {
	UserIds []int
	targetStatus int8
}

type UpdateUserStatusReq struct {
	UserIds []int
	TargetStatus int8
}