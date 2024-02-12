/*==============================================================*/
/* Nom de SGBD :  Sybase SQL Anywhere 11                        */
/* Date de cr�ation :  31/01/2024 15:09:16                      */
/*==============================================================*/


if exists(select 1 from sys.sysforeignkey where role='FK_COMMENT_AVOIR_POST') then
    alter table "Comment"
       delete foreign key FK_COMMENT_AVOIR_POST
end if;

if exists(select 1 from sys.sysforeignkey where role='FK_COMMENT_FAIRE_USER') then
    alter table "Comment"
       delete foreign key FK_COMMENT_FAIRE_USER
end if;

if exists(select 1 from sys.sysforeignkey where role='FK_EVENT_AVOIR_GROUP') then
    alter table Event
       delete foreign key FK_EVENT_AVOIR_GROUP
end if;

if exists(select 1 from sys.sysforeignkey where role='FK_EVENT_CREER_USER') then
    alter table Event
       delete foreign key FK_EVENT_CREER_USER
end if;

if exists(select 1 from sys.sysforeignkey where role='FK_FOLLOW_ASSOCIATI_USER') then
    alter table Follow
       delete foreign key FK_FOLLOW_ASSOCIATI_USER
end if;

if exists(select 1 from sys.sysforeignkey where role='FK_GROUP_CREER_USER') then
    alter table "Group"
       delete foreign key FK_GROUP_CREER_USER
end if;

if exists(select 1 from sys.sysforeignkey where role='FK_GROUPMES_ASSOCIATI_GROUP') then
    alter table GroupMessage
       delete foreign key FK_GROUPMES_ASSOCIATI_GROUP
end if;

if exists(select 1 from sys.sysforeignkey where role='FK_GROUPMES_MESAGEGRO_USER') then
    alter table GroupMessage
       delete foreign key FK_GROUPMES_MESAGEGRO_USER
end if;

if exists(select 1 from sys.sysforeignkey where role='FK_JOINNER_ASSOCIATI_GROUP') then
    alter table Joinner
       delete foreign key FK_JOINNER_ASSOCIATI_GROUP
end if;

if exists(select 1 from sys.sysforeignkey where role='FK_JOINNER_JOIN_USER') then
    alter table Joinner
       delete foreign key FK_JOINNER_JOIN_USER
end if;

if exists(select 1 from sys.sysforeignkey where role='FK_MESSAGE_ASSOCIATI_FOLLOW') then
    alter table "Message"
       delete foreign key FK_MESSAGE_ASSOCIATI_FOLLOW
end if;

if exists(select 1 from sys.sysforeignkey where role='FK_NOTFICAT_AVOIR_USER') then
    alter table Notfication
       delete foreign key FK_NOTFICAT_AVOIR_USER
end if;

if exists(select 1 from sys.sysforeignkey where role='FK_OPTION_ASSOCIATI_USER') then
    alter table "Option"
       delete foreign key FK_OPTION_ASSOCIATI_USER
end if;

if exists(select 1 from sys.sysforeignkey where role='FK_OPTION_AVOIR_EVENT') then
    alter table "Option"
       delete foreign key FK_OPTION_AVOIR_EVENT
end if;

if exists(select 1 from sys.sysforeignkey where role='FK_POST_ASSOCIATI_GROUP') then
    alter table Post
       delete foreign key FK_POST_ASSOCIATI_GROUP
end if;

if exists(select 1 from sys.sysforeignkey where role='FK_POST_CREATE_USER') then
    alter table Post
       delete foreign key FK_POST_CREATE_USER
end if;

if exists(select 1 from sys.sysforeignkey where role='FK_SESSION_ASSOCIATI_USER') then
    alter table "Session"
       delete foreign key FK_SESSION_ASSOCIATI_USER
end if;

if exists(
   select 1 from sys.sysindex i, sys.systable t
   where i.table_id=t.table_id 
     and i.index_name='ALLOWEDPOST_PK'
     and t.table_name='AllowedPost'
) then
   drop index AllowedPost.ALLOWEDPOST_PK
end if;

if exists(
   select 1 from sys.systable 
   where table_name='AllowedPost'
     and table_type in ('BASE', 'GBL TEMP')
) then
    drop table AllowedPost
end if;

if exists(
   select 1 from sys.sysindex i, sys.systable t
   where i.table_id=t.table_id 
     and i.index_name='FAIRE_FK'
     and t.table_name='Comment'
) then
   drop index "Comment".FAIRE_FK
end if;

if exists(
   select 1 from sys.sysindex i, sys.systable t
   where i.table_id=t.table_id 
     and i.index_name='AVOIR_FK'
     and t.table_name='Comment'
) then
   drop index "Comment".AVOIR_FK
end if;

if exists(
   select 1 from sys.sysindex i, sys.systable t
   where i.table_id=t.table_id 
     and i.index_name='COMMENT_PK'
     and t.table_name='Comment'
) then
   drop index "Comment".COMMENT_PK
end if;

if exists(
   select 1 from sys.systable 
   where table_name='Comment'
     and table_type in ('BASE', 'GBL TEMP')
) then
    drop table "Comment"
end if;

if exists(
   select 1 from sys.sysindex i, sys.systable t
   where i.table_id=t.table_id 
     and i.index_name='AVOIR_FK'
     and t.table_name='Event'
) then
   drop index Event.AVOIR_FK
end if;

if exists(
   select 1 from sys.sysindex i, sys.systable t
   where i.table_id=t.table_id 
     and i.index_name='CREER_FK'
     and t.table_name='Event'
) then
   drop index Event.CREER_FK
end if;

if exists(
   select 1 from sys.sysindex i, sys.systable t
   where i.table_id=t.table_id 
     and i.index_name='EVENT_PK'
     and t.table_name='Event'
) then
   drop index Event.EVENT_PK
end if;

if exists(
   select 1 from sys.systable 
   where table_name='Event'
     and table_type in ('BASE', 'GBL TEMP')
) then
    drop table Event
end if;

if exists(
   select 1 from sys.sysindex i, sys.systable t
   where i.table_id=t.table_id 
     and i.index_name='ASSOCIATION10_FK'
     and t.table_name='Follow'
) then
   drop index Follow.ASSOCIATION10_FK
end if;

if exists(
   select 1 from sys.sysindex i, sys.systable t
   where i.table_id=t.table_id 
     and i.index_name='FOLLOW_PK'
     and t.table_name='Follow'
) then
   drop index Follow.FOLLOW_PK
end if;

if exists(
   select 1 from sys.systable 
   where table_name='Follow'
     and table_type in ('BASE', 'GBL TEMP')
) then
    drop table Follow
end if;

if exists(
   select 1 from sys.sysindex i, sys.systable t
   where i.table_id=t.table_id 
     and i.index_name='CREER_FK'
     and t.table_name='Group'
) then
   drop index "Group".CREER_FK
end if;

if exists(
   select 1 from sys.sysindex i, sys.systable t
   where i.table_id=t.table_id 
     and i.index_name='GROUP_PK'
     and t.table_name='Group'
) then
   drop index "Group".GROUP_PK
end if;

if exists(
   select 1 from sys.systable 
   where table_name='Group'
     and table_type in ('BASE', 'GBL TEMP')
) then
    drop table "Group"
end if;

if exists(
   select 1 from sys.sysindex i, sys.systable t
   where i.table_id=t.table_id 
     and i.index_name='MESAGEGROUPE_FK'
     and t.table_name='GroupMessage'
) then
   drop index GroupMessage.MESAGEGROUPE_FK
end if;

if exists(
   select 1 from sys.sysindex i, sys.systable t
   where i.table_id=t.table_id 
     and i.index_name='ASSOCIATION6_FK'
     and t.table_name='GroupMessage'
) then
   drop index GroupMessage.ASSOCIATION6_FK
end if;

if exists(
   select 1 from sys.sysindex i, sys.systable t
   where i.table_id=t.table_id 
     and i.index_name='GROUPMESSAGE_PK'
     and t.table_name='GroupMessage'
) then
   drop index GroupMessage.GROUPMESSAGE_PK
end if;

if exists(
   select 1 from sys.systable 
   where table_name='GroupMessage'
     and table_type in ('BASE', 'GBL TEMP')
) then
    drop table GroupMessage
end if;

if exists(
   select 1 from sys.sysindex i, sys.systable t
   where i.table_id=t.table_id 
     and i.index_name='JOIN_FK'
     and t.table_name='Joinner'
) then
   drop index Joinner.JOIN_FK
end if;

if exists(
   select 1 from sys.sysindex i, sys.systable t
   where i.table_id=t.table_id 
     and i.index_name='ASSOCIATION17_FK'
     and t.table_name='Joinner'
) then
   drop index Joinner.ASSOCIATION17_FK
end if;

if exists(
   select 1 from sys.sysindex i, sys.systable t
   where i.table_id=t.table_id 
     and i.index_name='JOINNER_PK'
     and t.table_name='Joinner'
) then
   drop index Joinner.JOINNER_PK
end if;

if exists(
   select 1 from sys.systable 
   where table_name='Joinner'
     and table_type in ('BASE', 'GBL TEMP')
) then
    drop table Joinner
end if;

if exists(
   select 1 from sys.sysindex i, sys.systable t
   where i.table_id=t.table_id 
     and i.index_name='ASSOCIATION15_FK'
     and t.table_name='Message'
) then
   drop index "Message".ASSOCIATION15_FK
end if;

if exists(
   select 1 from sys.sysindex i, sys.systable t
   where i.table_id=t.table_id 
     and i.index_name='MESSAGE_PK'
     and t.table_name='Message'
) then
   drop index "Message".MESSAGE_PK
end if;

if exists(
   select 1 from sys.systable 
   where table_name='Message'
     and table_type in ('BASE', 'GBL TEMP')
) then
    drop table "Message"
end if;

if exists(
   select 1 from sys.sysindex i, sys.systable t
   where i.table_id=t.table_id 
     and i.index_name='NOTFICATION_AK'
     and t.table_name='Notfication'
) then
   drop index Notfication.NOTFICATION_AK
end if;

if exists(
   select 1 from sys.sysindex i, sys.systable t
   where i.table_id=t.table_id 
     and i.index_name='AVOIR_FK'
     and t.table_name='Notfication'
) then
   drop index Notfication.AVOIR_FK
end if;

if exists(
   select 1 from sys.systable 
   where table_name='Notfication'
     and table_type in ('BASE', 'GBL TEMP')
) then
    drop table Notfication
end if;

if exists(
   select 1 from sys.sysindex i, sys.systable t
   where i.table_id=t.table_id 
     and i.index_name='ASSOCIATION13_FK'
     and t.table_name='Option'
) then
   drop index "Option".ASSOCIATION13_FK
end if;

if exists(
   select 1 from sys.sysindex i, sys.systable t
   where i.table_id=t.table_id 
     and i.index_name='AVOIR_FK'
     and t.table_name='Option'
) then
   drop index "Option".AVOIR_FK
end if;

if exists(
   select 1 from sys.sysindex i, sys.systable t
   where i.table_id=t.table_id 
     and i.index_name='OPTION_PK'
     and t.table_name='Option'
) then
   drop index "Option".OPTION_PK
end if;

if exists(
   select 1 from sys.systable 
   where table_name='Option'
     and table_type in ('BASE', 'GBL TEMP')
) then
    drop table "Option"
end if;

if exists(
   select 1 from sys.sysindex i, sys.systable t
   where i.table_id=t.table_id 
     and i.index_name='ASSOCIATION16_FK'
     and t.table_name='Post'
) then
   drop index Post.ASSOCIATION16_FK
end if;

if exists(
   select 1 from sys.sysindex i, sys.systable t
   where i.table_id=t.table_id 
     and i.index_name='CREATE_FK'
     and t.table_name='Post'
) then
   drop index Post.CREATE_FK
end if;

if exists(
   select 1 from sys.sysindex i, sys.systable t
   where i.table_id=t.table_id 
     and i.index_name='POST_PK'
     and t.table_name='Post'
) then
   drop index Post.POST_PK
end if;

if exists(
   select 1 from sys.systable 
   where table_name='Post'
     and table_type in ('BASE', 'GBL TEMP')
) then
    drop table Post
end if;

if exists(
   select 1 from sys.sysindex i, sys.systable t
   where i.table_id=t.table_id 
     and i.index_name='ASSOCIATION9_FK'
     and t.table_name='Session'
) then
   drop index "Session".ASSOCIATION9_FK
end if;

if exists(
   select 1 from sys.sysindex i, sys.systable t
   where i.table_id=t.table_id 
     and i.index_name='SESSION_PK'
     and t.table_name='Session'
) then
   drop index "Session".SESSION_PK
end if;

if exists(
   select 1 from sys.systable 
   where table_name='Session'
     and table_type in ('BASE', 'GBL TEMP')
) then
    drop table "Session"
end if;

if exists(
   select 1 from sys.sysindex i, sys.systable t
   where i.table_id=t.table_id 
     and i.index_name='USER_PK'
     and t.table_name='User'
) then
   drop index "User".USER_PK
end if;

if exists(
   select 1 from sys.systable 
   where table_name='User'
     and table_type in ('BASE', 'GBL TEMP')
) then
    drop table "User"
end if;

/*==============================================================*/
/* Table : AllowedPost                                          */
/*==============================================================*/
create table AllowedPost 
(
   id                   integer                        not null,
   type                 integer                        null,
   constraint PK_ALLOWEDPOST primary key (id)
);

/*==============================================================*/
/* Index : ALLOWEDPOST_PK                                       */
/*==============================================================*/
create unique index ALLOWEDPOST_PK on AllowedPost (
id ASC
);

/*==============================================================*/
/* Table : "Comment"                                            */
/*==============================================================*/
create table "Comment" 
(
   id                   integer                        not null,
   Pos_id               integer                        not null,
   Use_id               integer                        not null,
   "comment"            varchar(254)                   null,
   type                 varchar(254)                   null,
   constraint PK_COMMENT primary key (id)
);

/*==============================================================*/
/* Index : COMMENT_PK                                           */
/*==============================================================*/
create unique index COMMENT_PK on "Comment" (
id ASC
);

/*==============================================================*/
/* Index : AVOIR_FK                                             */
/*==============================================================*/
create index AVOIR_FK on "Comment" (
Pos_id ASC
);

/*==============================================================*/
/* Index : FAIRE_FK                                             */
/*==============================================================*/
create index FAIRE_FK on "Comment" (
Use_id ASC
);

/*==============================================================*/
/* Table : Event                                                */
/*==============================================================*/
create table Event 
(
   id                   integer                        not null,
   Use_id               integer                        not null,
   Gro_id               integer                        not null,
   title                varchar(254)                   null,
   description          varchar(254)                   null,
   "date"               timestamp                      null,
   "time"               timestamp                      null,
   constraint PK_EVENT primary key (id)
);

/*==============================================================*/
/* Index : EVENT_PK                                             */
/*==============================================================*/
create unique index EVENT_PK on Event (
id ASC
);

/*==============================================================*/
/* Index : CREER_FK                                             */
/*==============================================================*/
create index CREER_FK on Event (
Use_id ASC
);

/*==============================================================*/
/* Index : AVOIR_FK                                             */
/*==============================================================*/
create index AVOIR_FK on Event (
Gro_id ASC
);

/*==============================================================*/
/* Table : Follow                                               */
/*==============================================================*/
create table Follow 
(
   id                   integer                        not null,
   Use_id               integer                        not null,
   isfollowing          smallint                       null,
   constraint PK_FOLLOW primary key (id)
);

/*==============================================================*/
/* Index : FOLLOW_PK                                            */
/*==============================================================*/
create unique index FOLLOW_PK on Follow (
id ASC
);

/*==============================================================*/
/* Index : ASSOCIATION10_FK                                     */
/*==============================================================*/
create index ASSOCIATION10_FK on Follow (
Use_id ASC
);

/*==============================================================*/
/* Table : "Group"                                              */
/*==============================================================*/
create table "Group" 
(
   id                   integer                        not null,
   Use_id               integer                        not null,
   titre                varchar(254)                   null,
   description          varchar(254)                   null,
   constraint PK_GROUP primary key (id)
);

/*==============================================================*/
/* Index : GROUP_PK                                             */
/*==============================================================*/
create unique index GROUP_PK on "Group" (
id ASC
);

/*==============================================================*/
/* Index : CREER_FK                                             */
/*==============================================================*/
create index CREER_FK on "Group" (
Use_id ASC
);

/*==============================================================*/
/* Table : GroupMessage                                         */
/*==============================================================*/
create table GroupMessage 
(
   id                   integer                        not null,
   Use_id               integer                        not null,
   Gro_id               integer                        not null,
   content              varchar(254)                   null,
   dateCreation         timestamp                      null,
   constraint PK_GROUPMESSAGE primary key (id)
);

/*==============================================================*/
/* Index : GROUPMESSAGE_PK                                      */
/*==============================================================*/
create unique index GROUPMESSAGE_PK on GroupMessage (
id ASC
);

/*==============================================================*/
/* Index : ASSOCIATION6_FK                                      */
/*==============================================================*/
create index ASSOCIATION6_FK on GroupMessage (
Gro_id ASC
);

/*==============================================================*/
/* Index : MESAGEGROUPE_FK                                      */
/*==============================================================*/
create index MESAGEGROUPE_FK on GroupMessage (
Use_id ASC
);

/*==============================================================*/
/* Table : Joinner                                              */
/*==============================================================*/
create table Joinner 
(
   id                   integer                        not null,
   Gro_id               integer                        not null,
   Use_id               integer                        not null,
   "date"               timestamp                      null,
   constraint PK_JOINNER primary key (id)
);

/*==============================================================*/
/* Index : JOINNER_PK                                           */
/*==============================================================*/
create unique index JOINNER_PK on Joinner (
id ASC
);

/*==============================================================*/
/* Index : ASSOCIATION17_FK                                     */
/*==============================================================*/
create index ASSOCIATION17_FK on Joinner (
Gro_id ASC
);

/*==============================================================*/
/* Index : JOIN_FK                                              */
/*==============================================================*/
create index JOIN_FK on Joinner (
Use_id ASC
);

/*==============================================================*/
/* Table : "Message"                                            */
/*==============================================================*/
create table "Message" 
(
   id                   integer                        not null,
   Fol_id               integer                        not null,
   content              varchar(254)                   null,
   dateCreation         timestamp                      null,
   constraint PK_MESSAGE primary key (id)
);

/*==============================================================*/
/* Index : MESSAGE_PK                                           */
/*==============================================================*/
create unique index MESSAGE_PK on "Message" (
id ASC
);

/*==============================================================*/
/* Index : ASSOCIATION15_FK                                     */
/*==============================================================*/
create index ASSOCIATION15_FK on "Message" (
Fol_id ASC
);

/*==============================================================*/
/* Table : Notfication                                          */
/*==============================================================*/
create table Notfication 
(
   Use_id               integer                        not null,
   Use_id               integer                        not null,
   id                   integer                        null,
   type                 varchar(254)                   null,
   status               varchar(254)                   null,
   constraint AK_NOTIFICATION_ID_NOTFICAT unique (id)
   constraint AK_NOTIFICATION_ID_NOTFICAT primary key (id)
);

CREATE TABLE notification (
    id INTEGER PRIMARY KEY,
    User_id INTEGER NOT NULL,
    Group_id INTEGER,
    type VARCHAR(255),
    status VARCHAR(255),
    FOREIGN KEY (User_id) REFERENCES "User" (id),
    FOREIGN KEY (Group_id) REFERENCES "Group" (id)
);


/*==============================================================*/
/* Index : AVOIR_FK                                             */
/*==============================================================*/
create index AVOIR_FK on Notfication (
Use_id ASC
);

/*==============================================================*/
/* Index : NOTFICATION_AK                                       */
/*==============================================================*/
create unique index NOTFICATION_AK on Notfication (
id ASC
);

/*==============================================================*/
/* Table : "Option"                                             */
/*==============================================================*/
create table "Option" 
(
   id                   integer                        not null,
   Eve_id               integer                        not null,
   Use_id               integer                        not null,
   isGoing              smallint                       null,
   constraint PK_OPTION primary key (id)
);

/*==============================================================*/
/* Index : OPTION_PK                                            */
/*==============================================================*/
create unique index OPTION_PK on "Option" (
id ASC
);

/*==============================================================*/
/* Index : AVOIR_FK                                             */
/*==============================================================*/
create index AVOIR_FK on "Option" (
Eve_id ASC
);

/*==============================================================*/
/* Index : ASSOCIATION13_FK                                     */
/*==============================================================*/
create index ASSOCIATION13_FK on "Option" (
Use_id ASC
);

/*==============================================================*/
/* Table : Post                                                 */
/*==============================================================*/
create table Post 
(
   id                   integer                        not null,
   Gro_id               integer                        null,
   Use_id               integer                        not null,
   titre                varchar(254)                   null,
   image                varchar(254)                   null,
   content              varchar(254)                   null,
   privaty              varchar(254)                   null,
   postGroup            varchar(254)                   null,
   creationDate         numeric                        null,
   constraint PK_POST primary key (id)
);

/*==============================================================*/
/* Index : POST_PK                                              */
/*==============================================================*/
create unique index POST_PK on Post (
id ASC
);

/*==============================================================*/
/* Index : CREATE_FK                                            */
/*==============================================================*/
create index CREATE_FK on Post (
Use_id ASC
);

/*==============================================================*/
/* Index : ASSOCIATION16_FK                                     */
/*==============================================================*/
create index ASSOCIATION16_FK on Post (
Gro_id ASC
);

/*==============================================================*/
/* Table : "Session"                                            */
/*==============================================================*/
create table "Session" 
(
   id                   integer                        not null,
   Use_id               integer                        not null,
   uiSession            integer                        null,
   email                varchar(254)                   null,
   constraint PK_SESSION primary key (id)
);

/*==============================================================*/
/* Index : SESSION_PK                                           */
/*==============================================================*/
create unique index SESSION_PK on "Session" (
id ASC
);

/*==============================================================*/
/* Index : ASSOCIATION9_FK                                      */
/*==============================================================*/
create index ASSOCIATION9_FK on "Session" (
Use_id ASC
);

/*==============================================================*/
/* Table : "User"                                               */
/*==============================================================*/
create table "User" 
(
   id                   integer                        not null,
   firstName            varchar(254)                   null,
   lastName             varchar(254)                   null,
   email                varchar(254)                   null,
   nickName             varchar(254)                   null,
   birthName            timestamp                      null,
   avatar               varchar(254)                   null,
   aboutMe              varchar(254)                   null,
   ispublic             smallint                       null,
   tokenLogin           varchar(254)                   null,
   constraint PK_USER primary key (id)
);

/*==============================================================*/
/* Index : USER_PK                                              */
/*==============================================================*/
create unique index USER_PK on "User" (
id ASC
);

alter table "Comment"
   add constraint FK_COMMENT_AVOIR_POST foreign key (Pos_id)
      references Post (id)
      on update restrict
      on delete restrict;

alter table "Comment"
   add constraint FK_COMMENT_FAIRE_USER foreign key (Use_id)
      references "User" (id)
      on update restrict
      on delete restrict;

alter table Event
   add constraint FK_EVENT_AVOIR_GROUP foreign key (Gro_id)
      references "Group" (id)
      on update restrict
      on delete restrict;

alter table Event
   add constraint FK_EVENT_CREER_USER foreign key (Use_id)
      references "User" (id)
      on update restrict
      on delete restrict;

alter table Follow
   add constraint FK_FOLLOW_ASSOCIATI_USER foreign key (Use_id)
      references "User" (id)
      on update restrict
      on delete restrict;

alter table "Group"
   add constraint FK_GROUP_CREER_USER foreign key (Use_id)
      references "User" (id)
      on update restrict
      on delete restrict;

alter table GroupMessage
   add constraint FK_GROUPMES_ASSOCIATI_GROUP foreign key (Gro_id)
      references "Group" (id)
      on update restrict
      on delete restrict;

alter table GroupMessage
   add constraint FK_GROUPMES_MESAGEGRO_USER foreign key (Use_id)
      references "User" (id)
      on update restrict
      on delete restrict;

alter table Joinner
   add constraint FK_JOINNER_ASSOCIATI_GROUP foreign key (Gro_id)
      references "Group" (id)
      on update restrict
      on delete restrict;

alter table Joinner
   add constraint FK_JOINNER_JOIN_USER foreign key (Use_id)
      references "User" (id)
      on update restrict
      on delete restrict;

alter table "Message"
   add constraint FK_MESSAGE_ASSOCIATI_FOLLOW foreign key (Fol_id)
      references Follow (id)
      on update restrict
      on delete restrict;

alter table Notfication
   add constraint FK_NOTFICAT_AVOIR_USER foreign key (Use_id)
      references "User" (id)
      on update restrict
      on delete restrict;

alter table "Option"
   add constraint FK_OPTION_ASSOCIATI_USER foreign key (Use_id)
      references "User" (id)
      on update restrict
      on delete restrict;

alter table "Option"
   add constraint FK_OPTION_AVOIR_EVENT foreign key (Eve_id)
      references Event (id)
      on update restrict
      on delete restrict;

alter table Post
   add constraint FK_POST_ASSOCIATI_GROUP foreign key (Gro_id)
      references "Group" (id)
      on update restrict
      on delete restrict;

alter table Post
   add constraint FK_POST_CREATE_USER foreign key (Use_id)
      references "User" (id)
      on update restrict
      on delete restrict;

alter table "Session"
   add constraint FK_SESSION_ASSOCIATI_USER foreign key (Use_id)
      references "User" (id)
      on update restrict
      on delete restrict;

