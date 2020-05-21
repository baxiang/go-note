package main

func GetMembersByLeaderId(id int64){
	select is_leader ,department_id from user_department where user_id= id
	//
	select id join user on department
}

func main() {
	
}
