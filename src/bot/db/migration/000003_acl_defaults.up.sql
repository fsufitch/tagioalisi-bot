-- add auto-incrementing default IDs for ACLs

ALTER TABLE user_acl ALTER COLUMN row_id SET DEFAULT nextval('seq_user_acl_row_id');
ALTER TABLE role_acl ALTER COLUMN row_id SET DEFAULT nextval('seq_role_acl_row_id');
