// Code generated by goctl. DO NOT EDIT.
package types

type UserLogin struct {
	TwitterUrl string `json:"twitterUrl"`
	InviteId   string `json:"inviteId"`
}

type UserLoginReply struct {
	Token string `json:"token"`
}

type AdminLogin struct {
	AdminName     string `json:"adminName"`
	AdminPassword string `json:"adminPassword"`
}

type AdminLoginReply struct {
	Token string `json:"token"`
}

type UserInfoReply struct {
	UserName    string `json:"userName"`
	TwitterName string `json:"twitterName"`
	UserAvatar  string `json:"userAvatar"`
	IsNew       int64  `json:"isNew"`
}

type HelpCategoryList struct {
	LanguageCode string `json:"languageCode"`
}

type HelpCategoryListReply struct {
	CategoryId   int64  `json:"categoryId"`
	CategoryName string `json:"categoryName"`
}

type CurrentOnlinePersonReply struct {
	CurrentOnlinePerson int64 `json:"currentOnlinePerson"`
}

type RegisteredPopulationReply struct {
	RegisteredPopulation int64 `json:"registeredPopulation"`
}

type GetUserListByCondition struct {
	BlackStatus int64  `json:"blackStatus"` //冻结状态
	PageNum     int64  `json:"pageNum"`     //页码
	PageSize    int64  `json:"pageSize"`    //每页条数
	OrderType   string `json:"order"`       //排序
	TwitterId   string `json:"twitterId"`   //推特id
}

type GetUserListByConditionReply struct {
	TwitterId     string `json:"twitterId"`     //推特id
	UserID        int64  `json:"userId"`        //用户id
	RegisterAt    int64  `json:"registerAt"`    //注册时间
	RecomendBy    string `json:"recomendBy"`    //推荐人
	IsBlacklisted int64  `json:"isBlacklisted"` //是否冻结
}

type BatchUpdateUserBlackStatus struct {
	UserIdList  []int64 `json:"userIdList"`  //用户id列表
	BlackStatus int64   `json:"blackStatus"` //冻结状态
}

type BatchUpdateUserBlackStatusReply struct {
	UserIdList  []int64 `json:"userIdList"`  //用户id列表
	BlackStatus int64   `json:"blackStatus"` //冻结状态
}
