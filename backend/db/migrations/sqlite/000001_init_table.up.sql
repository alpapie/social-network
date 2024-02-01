
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
   Post_id               integer                        not null,
   User_id               integer                        not null,
   "comment"            varchar(254)                   null,
   type                 varchar(254)                   null,
   constraint PK_COMMENT primary key (id),
   FOREIGN KEY (Post_id) REFERENCES Post (id) ON DELETE SET NULL,
   FOREIGN KEY (User_id) REFERENCES User (id) ON DELETE SET NULL
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
create index AVOIR_COMMENT_FK on "Comment" (
Post_id ASC
);

/*==============================================================*/
/* Index : FAIRE_FK                                             */
/*==============================================================*/
create index FAIRE_FK on "Comment" (
User_id ASC
);

/*==============================================================*/
/* Table : Event                                                */
/*==============================================================*/
create table Event 
(
   id                   integer                        not null,
   User_id               integer                        not null,
   Group_id               integer                        not null,
   title                varchar(254)                   null,
   description          varchar(254)                   null,
   "date"               timestamp                      null,
   "time"               timestamp                      null,
   constraint PK_EVENT primary key (id),
   FOREIGN KEY (Group_id) REFERENCES "Group" (id) ON DELETE SET NULL,
   FOREIGN KEY (User_id) REFERENCES User (id) ON DELETE SET NULL
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
create index CREER_EVENT_FK on Event (
User_id ASC
);

/*==============================================================*/
/* Index : AVOIR_FK                                             */
/*==============================================================*/
create index AVOIR___FK on Event (
Group_id ASC
);

/*==============================================================*/
/* Table : Follow                                               */
/*==============================================================*/
create table Follow 
(
   id                   integer                        not null,
   User_id               integer                        not null,
   isfollowing          smallint                       null,
   constraint PK_FOLLOW primary key (id),
   FOREIGN KEY (User_id) REFERENCES User (id) ON DELETE SET NULL
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
User_id ASC
);

/*==============================================================*/
/* Table : "Group"                                              */
/*==============================================================*/
create table "Group" 
(
   id                   integer                        not null,
   User_id               integer                        not null,
   titre                varchar(254)                   null,
   description          varchar(254)                   null,
   constraint PK_GROUP primary key (id),
   FOREIGN KEY (User_id) REFERENCES User (id) ON DELETE SET NULL
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
create index CREER__FK on "Group" (
User_id ASC
);

/*==============================================================*/
/* Table : GroupMessage                                         */
/*==============================================================*/
create table GroupMessage 
(
   id                   integer                        not null,
   User_id               integer                        not null,
   Group_id               integer                        not null,
   content              varchar(254)                   null,
   dateCreation         timestamp                      null,
   constraint PK_GROUPMESSAGE primary key (id),
   FOREIGN KEY (Group_id) REFERENCES "Group" (id) ON DELETE SET NULL,
   FOREIGN KEY (User_id) REFERENCES User (id) ON DELETE SET NULL
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
Group_id ASC
);

/*==============================================================*/
/* Index : MESAGEGROUPE_FK                                      */
/*==============================================================*/
create index MESAGEGROUPE_FK on GroupMessage (
User_id ASC
);

/*==============================================================*/
/* Table : Joinner                                              */
/*==============================================================*/
create table Joinner 
(
   id                   integer                        not null,
   Group_id               integer                        not null,
   User_id               integer                        not null,
   "date"               timestamp                      null,
   constraint PK_JOINNER primary key (id),
   FOREIGN KEY (Group_id) REFERENCES "Group" (id) ON DELETE SET NULL,
   FOREIGN KEY (User_id) REFERENCES User (id) ON DELETE SET NULL

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
Group_id ASC
);

/*==============================================================*/
/* Index : JOIN_FK                                              */
/*==============================================================*/
create index JOIN_FK on Joinner (
User_id ASC
);

/*==============================================================*/
/* Table : "Message"                                            */
/*==============================================================*/
create table "Message" 
(
   id                   integer                        not null,
   Follow_id               integer                        not null,
   content              varchar(254)                   null,
   dateCreation         timestamp                      null,
   constraint PK_MESSAGE primary key (id),
   FOREIGN KEY (Follow_id) REFERENCES Follow (id) ON DELETE SET NULL
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
Follow_id ASC
);

/*==============================================================*/
/* Table : Notfication                                          */
/*==============================================================*/
create table Notfication 
(
   User_id               integer                        not null,
   id                   integer                        null,
   type                 varchar(254)                   null,
   status               varchar(254)                   null,
   constraint AK_NOTIFICATION_ID_NOTFICAT unique (id),
   FOREIGN KEY (User_id) REFERENCES User (id) ON DELETE SET NULL
);

/*==============================================================*/
/* Index : AVOIR_FK                                             */
/*==============================================================*/
create index AVOIRNOT_FK on Notfication (
User_id ASC
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
   User_id               integer                        not null,
   isGoing              smallint                       null,
   constraint PK_OPTION primary key (id),
   FOREIGN KEY (User_id) REFERENCES User (id) ON DELETE SET NULL,
   FOREIGN KEY (Eve_id) REFERENCES Event (id) ON DELETE SET NULL
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
create index AVOIRPTIO_FK on "Option" (
Eve_id ASC
);

/*==============================================================*/
/* Index : ASSOCIATION13_FK                                     */
/*==============================================================*/
create index ASSOCIATION13_FK on "Option" (
User_id ASC
);

/*==============================================================*/
/* Table : Post                                                 */
/*==============================================================*/
create table Post 
( 
   id                   integer                        not null,
   Group_id               integer                        null,
   User_id               integer                        not null,
   titre                varchar(254)                   null,
   image                varchar(254)                   null,
   content              varchar(254)                   null,
   privacy              varchar(254)    CHECK (primary IN ('public', 'private' , 'almostprivate'))   null,
   postGroup            varchar(254)                   null,
   creationDate         numeric                         null,
   constraint PK_POST primary key (id),
   FOREIGN KEY (User_id) REFERENCES User (id) ON DELETE SET NULL,
   FOREIGN KEY (Group_id) REFERENCES "Group" (id) ON DELETE SET NULL
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
User_id ASC
);

/*==============================================================*/
/* Index : ASSOCIATION16_FK                                     */
/*==============================================================*/
create index ASSOCIATION16_FK on Post (
Group_id ASC
);

/*==============================================================*/
/* Table : "Session"                                            */
/*==============================================================*/
create table "Session" 
(
   id                   integer                        not null,
   User_id               integer                        not null,
   uiSession            integer                        null,
   email                varchar(254)                   null,
   constraint PK_SESSION primary key (id),
   FOREIGN KEY (User_id) REFERENCES User (id) ON DELETE SET NULL
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
User_id ASC
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
