create table roles_permissions(roles_id int, permission_id int, CONSTRAINT PRIMARY KEY(roles_id, permission_id), FOREIGN KEY (roles_id) REFERENCES roles(id), FOREIGN KEY (permission_id) REFERENCES  permissions(id));