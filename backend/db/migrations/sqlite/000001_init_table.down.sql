-- SQLBook: Code


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
