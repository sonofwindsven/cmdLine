create table `blog_tag` (
	`id` int(10) UNSIGNED not null auto_increment,
	`name` varchar(100) DEFAULT '' comment '标签名称',
	`create_on` int(10) UNSIGNED DEFAULT 0 comment '创建时间',
  `create_by` varchar(100) default '' comment '创建人',
  `modified_on` int(10) UNSIGNED DEFAULT 0 comment '修改时间',
  `modified_by` varchar(100) default '' comment '修改人' ,
  `deleted_on` int(10) UNSIGNED DEFAULT 0 comment '删除时间',
  `is_del` tinyint(3) UNSIGNED default 0 comment '是否删除 0:未删除 1:已删除',
  `state` tinyint(3) UNSIGNED default 1 comment '状态 0:禁用 1:启用',
  primary key(`id`)
) engine=InnoDB DEFAULT charset=utf8mb4 comment='标签管理';
