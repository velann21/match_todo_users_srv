release:
	docker build -t "singaravelan21/match_todo_list_user_srv" .;
	docker tag "singaravelan21/match_todo_list_user_srv" "singaravelan21/match_todo_list_user_srv:v2.17.0";
	docker push "singaravelan21/match_todo_list_user_srv:v2.17.0"