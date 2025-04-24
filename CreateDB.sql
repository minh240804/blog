-- drop table "Comment";
-- drop table blog;
-- drop table "User";
-- drop table category;

create table Category(
	CategoryId int primary key generated always as identity (increment by 1),
	categoryName varchar(255) not null,
	description varchar(512) not null,  
	createAt timestamp with time zone default NOW(),
	updateAt timestamp with time zone default NOW(),
	delateAt timestamp with time zone
);

create table "User" (
	userId int primary key generated always as identity (increment by 1),
	UserName Varchar(255) not null,
	"Password" Varchar(255) not null,
	Role Varchar(255) default 'customer',
	createAt timestamp with time zone default NOW(),
	updateAt timestamp with time zone default NOW(),
	delateAt timestamp with time zone
);

create table Blog(
	blogID int primary key generated always as identity (increment by 1),
	blogName Varchar(512) not null,
	blogContent Varchar(10197) not null,
	vote int,
	categoryId int ,
	authorId int,

	createAt timestamp with time zone default NOW(),
	updateAt timestamp with time zone default NOW(),
	delateAt timestamp with time zone
);

create table "Comment"(
	commentId int primary key generated always as identity (increment by 1),
	blogId int not null,
	authorId int not  null,
	"content" varchar(512),

	createAt timestamp with time zone default NOW(),
	updateAt timestamp with time zone default NOW(),
	delateAt timestamp with time zone
);

Alter table blog
	add constraint fk_topic foreign Key (categoryId) REFERENCES Category (CategoryId),
	add constraint fk_write foreign Key (authorId) REFERENCES "User" (userId);

Alter table "Comment"
	add constraint fk_comment foreign Key (blogid) REFERENCES Blog (blogID),
	add constraint fk_write foreign Key (authorid) REFERENCES "User" (userId);