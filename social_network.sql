/*==============================================================*/
/* Nom de SGBD :  Sybase SQL Anywhere 11                        */
/* Date de création :  30/01/2024 16:43:02                      */
/*==============================================================*/


if exists(select 1 from sys.sysforeignkey where role='FK_ALLOWEDU_CONSULTER_POST') then
    alter table AllowedUser
       delete foreign key FK_ALLOWEDU_CONSULTER_POST
end if;

if exists(select 1 from sys.sysforeignkey where role='FK_CHAT_ASSOCIATI_FOLLOW') then
    alter table Chat
       delete foreign key FK_CHAT_ASSOCIATI_FOLLOW
end if;

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

if exists(select 1 from sys.sysforeignkey where role='FK_FOLLOW_ASSOCIATI_ALLOWEDU') then
    alter table Follow
       delete foreign key FK_FOLLOW_ASSOCIATI_ALLOWEDU
end if;

if exists(select 1 from sys.sysforeignkey where role='FK_GROUP_ASSOCIATI_CHAT') then
    alter table "Group"
       delete foreign key FK_GROUP_ASSOCIATI_CHAT
end if;

if exists(select 1 from sys.sysforeignkey where role='FK_GROUP_CREER_USER') then
    alter table "Group"
       delete foreign key FK_GROUP_CREER_USER
end if;

if exists(select 1 from sys.sysforeignkey where role='FK_MESSAGE_ASSOCIATI_CHAT') then
    alter table "Message"
       delete foreign key FK_MESSAGE_ASSOCIATI_CHAT
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

if exists(select 1 from sys.sysforeignkey where role='FK_POST_CREATE_USER') then
    alter table Post
       delete foreign key FK_POST_CREATE_USER
end if;

if exists(select 1 from sys.sysforeignkey where role='FK_SESSION_ASSOCIATI_USER') then
    alter table "Session"
       delete foreign key FK_SESSION_ASSOCIATI_USER
end if;

if exists(select 1 from sys.sysforeignkey where role='FK_CHATTER_CHATTER_CHAT') then
    alter table chatter
       delete foreign key FK_CHATTER_CHATTER_CHAT
end if;

if exists(select 1 from sys.sysforeignkey where role='FK_CHATTER_CHATTER_USER') then
    alter table chatter
       delete foreign key FK_CHATTER_CHATTER_USER
end if;

if exists(
   select 1 from sys.sysindex i, sys.systable t
   where i.table_id=t.table_id 
     and i.index_name='CONSULTER_FK'
     and t.table_name='AllowedUser'
) then
   drop index AllowedUser.CONSULTER_FK
end if;

if exists(
   select 1 from sys.sysindex i, sys.systable t
   where i.table_id=t.table_id 
     and i.index_name='ALLOWEDUSER_PK'
     and t.table_name='AllowedUser'
) then
   drop index AllowedUser.ALLOWEDUSER_PK
end if;

if exists(
   select 1 from sys.systable 
   where table_name='AllowedUser'
     and table_type in ('BASE', 'GBL TEMP')
) then
    drop table AllowedUser
end if;

if exists(
   select 1 from sys.sysindex i, sys.systable t
   where i.table_id=t.table_id 
     and i.index_name='ASSOCIATION15_FK'
     and t.table_name='Chat'
) then
   drop index Chat.ASSOCIATION15_FK
end if;

if exists(
   select 1 from sys.sysindex i, sys.systable t
   where i.table_id=t.table_id 
     and i.index_name='CHAT_PK'
     and t.table_name='Chat'
) then
   drop index Chat.CHAT_PK
end if;

if exists(
   select 1 from sys.systable 
   where table_name='Chat'
     and table_type in ('BASE', 'GBL TEMP')
) then
    drop table Chat
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
     and i.index_name='ASSOCIATION11_FK'
     and t.table_name='Follow'
) then
   drop index Follow.ASSOCIATION11_FK
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
     and i.index_name='ASSOCIATION6_FK'
     and t.table_name='Group'
) then
   drop index "Group".ASSOCIATION6_FK
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
     and i.index_name='ASSOCIATION7_FK'
     and t.table_name='Message'
) then
   drop index "Message".ASSOCIATION7_FK
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

if exists(
   select 1 from sys.sysindex i, sys.systable t
   where i.table_id=t.table_id 
     and i.index_name='CHATTER_FK2'
     and t.table_name='chatter'
) then
   drop index chatter.CHATTER_FK2
end if;

if exists(
   select 1 from sys.sysindex i, sys.systable t
   where i.table_id=t.table_id 
     and i.index_name='CHATTER_FK'
     and t.table_name='chatter'
) then
   drop index chatter.CHATTER_FK
end if;

if exists(
   select 1 from sys.sysindex i, sys.systable t
   where i.table_id=t.table_id 
     and i.index_name='CHATTER_PK'
     and t.table_name='chatter'
) then
   drop index chatter.CHATTER_PK
end if;

if exists(
   select 1 from sys.systable 
   where table_name='chatter'
     and table_type in ('BASE', 'GBL TEMP')
) then
    drop table chatter
end if;

/*==============================================================*/
/* Table : AllowedUser                                          */
/*==============================================================*/
create table AllowedUser 
(
   id                   integer                        not null,
   Pos_id               integer                        not null,
   type                 integer                        null,
   constraint PK_ALLOWEDUSER primary key (id)
);

/*==============================================================*/
/* Index : ALLOWEDUSER_PK                                       */
/*==============================================================*/
create unique index ALLOWEDUSER_PK on AllowedUser (
id ASC
);

/*==============================================================*/
/* Index : CONSULTER_FK                                         */
/*==============================================================*/
create index CONSULTER_FK on AllowedUser (
Pos_id ASC
);

/*==============================================================*/
/* Table : Chat                                                 */
/*==============================================================*/
create table Chat 
(
   id                   integer                        not null,
   Fol_id               integer                        not null,
   "update"             timestamp                      null,
   constraint PK_CHAT primary key (id)
);

/*==============================================================*/
/* Index : CHAT_PK                                              */
/*==============================================================*/
create unique index CHAT_PK on Chat (
id ASC
);

/*==============================================================*/
/* Index : ASSOCIATION15_FK                                     */
/*==============================================================*/
create index ASSOCIATION15_FK on Chat (
Fol_id ASC
);

/*==============================================================*/
/* Table : "Comment"                                            */
/*==============================================================*/
create table "Comment" 
(
   id                   integer                        not null,
   Use_id               integer                        not null,
   Pos_id               integer                        not null,
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
   All_id               integer                        not null,
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
/* Index : ASSOCIATION11_FK                                     */
/*==============================================================*/
create index ASSOCIATION11_FK on Follow (
All_id ASC
);

/*==============================================================*/
/* Table : "Group"                                              */
/*==============================================================*/
create table "Group" 
(
   id                   integer                        not null,
   Use_id               integer                        not null,
   Cha_id               integer                        null,
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
/* Index : ASSOCIATION6_FK                                      */
/*==============================================================*/
create index ASSOCIATION6_FK on "Group" (
Cha_id ASC
);

/*==============================================================*/
/* Index : CREER_FK                                             */
/*==============================================================*/
create index CREER_FK on "Group" (
Use_id ASC
);

/*==============================================================*/
/* Table : "Message"                                            */
/*==============================================================*/
create table "Message" 
(
   id                   integer                        not null,
   Cha_id               integer                        not null,
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
/* Index : ASSOCIATION7_FK                                      */
/*==============================================================*/
create index ASSOCIATION7_FK on "Message" (
Cha_id ASC
);

/*==============================================================*/
/* Table : Notfication                                          */
/*==============================================================*/
create table Notfication 
(
   Use_id               integer                        not null,
   id                   integer                        null,
   type                 varchar(254)                   null,
   status               varchar(254)                   null,
   constraint AK_NOTIFICATION_ID_NOTFICAT unique (id)
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
   Use_id               integer                        not null,
   Eve_id               integer                        not null,
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

/*==============================================================*/
/* Table : chatter                                              */
/*==============================================================*/
create table chatter 
(
   Use_id               integer                        not null,
   id                   integer                        not null,
   constraint PK_CHATTER primary key clustered (Use_id, id)
);

/*==============================================================*/
/* Index : CHATTER_PK                                           */
/*==============================================================*/
create unique clustered index CHATTER_PK on chatter (
Use_id ASC,
id ASC
);

/*==============================================================*/
/* Index : CHATTER_FK                                           */
/*==============================================================*/
create index CHATTER_FK on chatter (
Use_id ASC
);

/*==============================================================*/
/* Index : CHATTER_FK2                                          */
/*==============================================================*/
create index CHATTER_FK2 on chatter (
id ASC
);

alter table AllowedUser
   add constraint FK_ALLOWEDU_CONSULTER_POST foreign key (Pos_id)
      references Post (id)
      on update restrict
      on delete restrict;

alter table Chat
   add constraint FK_CHAT_ASSOCIATI_FOLLOW foreign key (Fol_id)
      references Follow (id)
      on update restrict
      on delete restrict;

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

alter table Follow
   add constraint FK_FOLLOW_ASSOCIATI_ALLOWEDU foreign key (All_id)
      references AllowedUser (id)
      on update restrict
      on delete restrict;

alter table "Group"
   add constraint FK_GROUP_ASSOCIATI_CHAT foreign key (Cha_id)
      references Chat (id)
      on update restrict
      on delete restrict;

alter table "Group"
   add constraint FK_GROUP_CREER_USER foreign key (Use_id)
      references "User" (id)
      on update restrict
      on delete restrict;

alter table "Message"
   add constraint FK_MESSAGE_ASSOCIATI_CHAT foreign key (Cha_id)
      references Chat (id)
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
   add constraint FK_POST_CREATE_USER foreign key (Use_id)
      references "User" (id)
      on update restrict
      on delete restrict;

alter table "Session"
   add constraint FK_SESSION_ASSOCIATI_USER foreign key (Use_id)
      references "User" (id)
      on update restrict
      on delete restrict;

alter table chatter
   add constraint FK_CHATTER_CHATTER_CHAT foreign key (id)
      references Chat (id)
      on update restrict
      on delete restrict;

alter table chatter
   add constraint FK_CHATTER_CHATTER_USER foreign key (Use_id)
      references "User" (id)
      on update restrict
      on delete restrict;

