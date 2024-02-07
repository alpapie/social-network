

-- ALTER TABLE AllowedPost ADD Post_id integer;
-- ALTER TABLE AllowedPost ADD User_id integer;
-- ALTER TABLE AllowedPost DROP COLUMN type;
-- ALTER TABLE AllowedPost DROP PRIMARY KEY;
-- ALTER TABLE AllowedPost ADD PRIMARY KEY (Post_id, User_id);
-- ALTER TABLE AllowedPost ADD CONSTRAINT  FOREIGN KEY (Post_id) REFERENCES Post(id);
-- ALTER TABLE AllowedPost ADD CONSTRAINT  FOREIGN KEY (User_id) REFERENCES User(id);