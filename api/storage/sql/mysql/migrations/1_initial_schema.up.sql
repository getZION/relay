
DROP TABLE IF EXISTS users;
CREATE TABLE users (
  id                bigint        NOT NULL AUTO_INCREMENT,
  did               varchar(256)  NOT NULL,
  username          varchar(16)   NOT NULL,
  email             varchar(64)   DEFAULT NULL,
  name              varchar(32)   NOT NULL,
  bio               varchar(512)  DEFAULT NULL,
  img               varchar(256)  DEFAULT NULL,
  price_to_message  bigint        DEFAULT NULL,
  created           bigint        NOT NULL,
  updated           bigint        DEFAULT NULL,
  PRIMARY KEY (id),
  UNIQUE KEY did (did),
  UNIQUE KEY id (id),
  UNIQUE KEY username (username)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;

DROP TABLE IF EXISTS communities;
CREATE TABLE communities (
  id                bigint        NOT NULL AUTO_INCREMENT,
  zid               varchar(256)  NOT NULL,
  name              varchar(128)  NOT NULL,
  owner_did         varchar(256)  NOT NULL,
  owner_username    varchar(16)   NOT NULL,
  description       varchar(512)  NOT NULL,
  escrow_amount     bigint        NOT NULL,
  img               varchar(512)  DEFAULT NULL,
  last_active       bigint        DEFAULT NULL,
  price_per_message bigint        NOT NULL,
  price_to_join     bigint        NOT NULL,
  public            boolean       DEFAULT 1,
  created           bigint        NOT NULL,
  updated           bigint        DEFAULT NULL,
  deleted           boolean       DEFAULT 0,
  PRIMARY KEY (id),
  UNIQUE KEY id (id),
  UNIQUE KEY name (name),
  UNIQUE KEY zid (zid),
  CONSTRAINT communities_user_did_users_Did_foreign FOREIGN KEY (owner_did) REFERENCES users (did) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;

DROP TABLE IF EXISTS conversations;
CREATE TABLE conversations (
  id            bigint        NOT NULL AUTO_INCREMENT,
  zid           varchar(256)  DEFAULT NULL,
  community_zid varchar(256)  NOT NULL,
  user_did      varchar(256)  NOT NULL,
  text          varchar(256)  DEFAULT NULL,
  link          varchar(256)  DEFAULT NULL,
  img           varchar(256)  DEFAULT NULL,
  video         varchar(256)  DEFAULT NULL,
  public        boolean       DEFAULT 1,
  public_price  bigint        DEFAULT '0',
  created       bigint        NOT NULL,
  updated       bigint        DEFAULT NULL,
  deleted       boolean       DEFAULT 0,
  PRIMARY KEY (id),
  UNIQUE KEY id (id),
  UNIQUE KEY zid (zid),
  KEY community_zid (community_zid),
  CONSTRAINT conversations_community_zid_communities_Zid_foreign FOREIGN KEY (community_zid) REFERENCES communities (zid) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT conversation_user_did_users_Did_foreign FOREIGN KEY (user_did) REFERENCES users (did) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;

DROP TABLE IF EXISTS tags;
CREATE TABLE tags (
  tag VARCHAR(64) NOT NULL PRIMARY KEY
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;

DROP TABLE IF EXISTS comments;
CREATE TABLE comments (
  id                bigint        NOT NULL AUTO_INCREMENT,
  zid               varchar(256)  NOT NULL,
  conversation_zid  varchar(256)  NOT NULL,
  user_did          varchar(256)  NOT NULL,
  text              varchar(256)  DEFAULT NULL,
  link              varchar(256)  DEFAULT NULL,
  created           bigint        NOT NULL,
  updated           bigint        DEFAULT NULL,
  deleted           boolean       DEFAULT 0,
  PRIMARY KEY (id),
  UNIQUE KEY id (id),
  KEY conversation_zid (conversation_zid),
  KEY user_did (user_did),
  CONSTRAINT comments_conversation_zid_conversations_Zid_foreign FOREIGN KEY (conversation_zid) REFERENCES conversations (zid) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT comments_user_did_users_Did_foreign FOREIGN KEY (user_did) REFERENCES users (did) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;

DROP TABLE IF EXISTS community_tags;
CREATE TABLE community_tags (
    community_zid   VARCHAR(256)   NOT NULL,
    tag             VARCHAR(64)    NOT NULL,
	  FOREIGN KEY (community_zid) REFERENCES  communities(zid),
    FOREIGN KEY (tag)           REFERENCES  tags(tag),
    UNIQUE  KEY community_zid_tag (community_zid, tag)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;

DROP TABLE IF EXISTS community_users;
CREATE TABLE community_users (
  id              bigint       NOT NULL AUTO_INCREMENT,
  community_zid   varchar(256) NOT NULL,
  user_did        varchar(256) NOT NULL,
  joined_date     bigint       NOT NULL,
  left_date       bigint       DEFAULT NULL,
  left_reason     varchar(64)  NULL,
  PRIMARY KEY (id),
  KEY user_did (user_did),
  KEY community_zid (community_zid),
  CONSTRAINT community_users_community_zid_communities_Zid_foreign FOREIGN KEY (community_zid) REFERENCES communities (zid) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT community_users_user_did_users_Did_foreign FOREIGN KEY (user_did) REFERENCES users (did) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci
