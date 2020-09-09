package controllers

import (
	//"github.com/coreos/etcd/clientv3"
	//"github.com/coreos/etcd/clientv3"
)

const USERKEY  = "todo/users/"
type UserController struct {
	//EtcdClient *clientv3.Client
}

//func (gctl *UserController) HealthyMe(res http.ResponseWriter, req *http.Request){
//	resp := SuccessResponse{}
//	resp.GeneralSuccessResp()
//	res.WriteHeader(200)
//	_ = json.NewEncoder(res).Encode(resp)
//}
//
//func (gctl *UserController) UsersRegistrationController(res http.ResponseWriter, req *http.Request){
//	user := User{}
//	resp := SuccessResponse{}
//	err := user.PopulateUserSigninRequest(req.Body)
//	if err != nil{
//		logrus.WithError(err).Error("failed to PopulateUserSigninRequest")
//		resp.GeneralResp()
//		res.WriteHeader(500)
//		_ = json.NewEncoder(res).Encode(resp)
//		return
//	}
//
//	d, err := json.Marshal(user)
//	if err != nil {
//		logrus.WithError(err).Error("failed to PopulateUserSigninRequest")
//		resp.GeneralResp()
//		res.WriteHeader(500)
//		_ = json.NewEncoder(res).Encode(resp)
//		return
//	}
//	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
//	defer cancel()
//	_, err = gctl.EtcdClient.Put(ctx, USERKEY + user.FirstName+user.LastName, string(d))
//	if err != nil{
//		logrus.WithError(err).Error("failed to EtcdClient PUT")
//		resp.GeneralResp()
//		res.WriteHeader(500)
//		_ = json.NewEncoder(res).Encode(resp)
//		return
//	}
//	resp.UserRegistrationResp(&user.FirstName)
//	res.WriteHeader(200)
//	_ = json.NewEncoder(res).Encode(resp)
//}
//
//func (gctl *UserController) ListUserController(res http.ResponseWriter, req *http.Request){
//	resp := SuccessResponse{}
//	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
//	defer cancel()
//	usersResp, err := gctl.EtcdClient.Get(ctx, USERKEY, clientv3.WithPrefix())
//	if err != nil{
//		logrus.WithError(err).Error("failed to EtcdClient Get")
//		resp.GeneralResp()
//		res.WriteHeader(500)
//		_ = json.NewEncoder(res).Encode(resp)
//		return
//	}
//	if len(usersResp.Kvs) <= 0 {
//		resp.GeneralResp()
//		res.WriteHeader(500)
//		_ = json.NewEncoder(res).Encode(resp)
//	}
//	var u User
//	err = json.Unmarshal(usersResp.Kvs[0].Value, &u)
//	if err != nil {
//		resp.GeneralResp()
//		res.WriteHeader(500)
//		_ = json.NewEncoder(res).Encode(resp)
//	}
//	resp.ListUserController(u)
//	res.WriteHeader(200)
//	_ = json.NewEncoder(res).Encode(resp)
//}
//
//func (gctl *UserController) ListUsersController(res http.ResponseWriter, req *http.Request){
//	resp := SuccessResponse{}
//	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
//	defer cancel()
//	usersResp, err := gctl.EtcdClient.Get(ctx, USERKEY, clientv3.WithPrefix())
//	if err != nil{
//		logrus.WithError(err).Error("failed to EtcdClient Get")
//		resp.GeneralResp()
//		res.WriteHeader(500)
//		_ = json.NewEncoder(res).Encode(resp)
//		return
//	}
//	if len(usersResp.Kvs) <= 0 {
//		resp.GeneralResp()
//		_ = json.NewEncoder(res).Encode(resp)
//	}
//	var u []User
//	for _ ,v := range usersResp.Kvs{
//		userObj := User{}
//		err = json.Unmarshal(v.Value, &userObj)
//		if err != nil{
//			continue
//		}
//		u = append(u, userObj)
//	}
//	resp.ListUsersController(u)
//	res.WriteHeader(200)
//	_ = json.NewEncoder(res).Encode(resp)
//}
//
//type User struct {
//	FirstName string
//	LastName  string
//	Company string
//	Password string
//}
//
//type Address struct {
//	Name string
//	Country string
//}
//
//func (userSignin *User) PopulateUserSigninRequest(body io.ReadCloser) error {
//	decoder := json.NewDecoder(body)
//	err := decoder.Decode(&userSignin)
//	if err != nil {
//		return errors.New("Invalid request")
//	}
//	return nil
//}
//
//
//
//type SuccessResponse struct {
//	Success bool
//	Status string
//	Data []map[string]interface{}
//}
//
//func (entity *SuccessResponse) UserRegistrationResp(id *string) {
//	responseData := make([]map[string]interface{}, 0)
//	data := make(map[string]interface{})
//	data["id"] = *id
//	responseData = append(responseData, data)
//	entity.Data = responseData
//	metaData := make(map[string]interface{})
//	metaData["message"] = "User registered"
//}
//
//func (entity *SuccessResponse) GeneralResp(){
//	entity.Success = false
//	entity.Status = "Error occured"
//	responseData := make([]map[string]interface{}, 0)
//	entity.Data = responseData
//}
//
//func (entity *SuccessResponse) GeneralSuccessResp(){
//	entity.Success = true
//	entity.Status = "Success"
//	responseData := make([]map[string]interface{}, 0)
//	entity.Data = responseData
//}
//
//func (entity *SuccessResponse) ListUsersController(users []User) {
//	responseData := make([]map[string]interface{}, 0)
//	data := make(map[string]interface{})
//	data["user"] = users
//	responseData = append(responseData, data)
//	entity.Data = responseData
//	metaData := make(map[string]interface{})
//	metaData["message"] = "User list"
//}
//
//func (entity *SuccessResponse) ListUserController(user User) {
//	responseData := make([]map[string]interface{}, 0)
//	data := make(map[string]interface{})
//	data["users"] = user
//	responseData = append(responseData, data)
//	entity.Data = responseData
//	metaData := make(map[string]interface{})
//	metaData["message"] = "User"
//}
//
//
//
//
//
