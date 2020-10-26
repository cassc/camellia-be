create table feedback(
id integer primary key autoincrement,
username text,
message text,
email text,
phone text
);

create table news(
id integer primary key autoincrement,
title text,
`source` text,
abstract text,
`date` text
);

insert into news(id, title,`source`,abstract, `date`) values
       (1, '高速发展的RFID，助力智能包装行业发展', 'RFID世界网', '中国智能包装行业快速发展为诸多传统包装印刷企业带来了新的发展机遇。未来随着印刷电子、RFID、柔性显示等创新技术发展与深度融合，尤其是RFID技术与电子标签的高速发展，将为智能包装及业内企业的发展带来利好。', '2019-03-13'),
       (2, '区块链如何赋能物联网', '波士顿咨询', '本文剖析区块链和物联网两种前沿技术结合后的应用趋势。研究表明，只有一小部分公司正在试验区块链与物联网技术结合解决方案，其中大多尚处于概念验证阶段。但这两种新兴技术双剑合璧，可在诸多应用场景和业务领域产生可持续的竞争优势。', '2019-02-13'),
       (3, '结合LoRa 及Cellular IoT产生互补性之新商机', 'stpi.nar', '借助跨平台连接进行布署类似LoRaWAN及蜂巢式 物联网多种技术的融合，如此一来，电信商不再需要倚靠物联网，而能透过单一平台便能解决手上所有的装置连接工作。', '2019-03-07'),
       (4, 'NFC与RFID技术的发展与创新', '搜狐', '当今时代，手机是人们日常生活的随身物品，它甚至逐渐取代了人们日常生活中的一些生活工具，如钥匙、门禁卡、信用卡、车票等等，而通过NFC技术，手机的功能更是日益强大。', '2019-03-13'),
       (5, '全球工业物联网市场生态全景观察', '资本实验室', '在当前，诸如5G，物联网、边缘计算、人工智能、机器人、区块链、增材制造和虚拟现实/增强现实等技术正在加速融合到工业物联网（Industrial Internet of Things，IIoT）的肥沃土壤中，并共同推动“第四次工业革命”或者说“工业4．0”的到来。', '2019-03-07'),
       (6, 'RFID射频识别系统详解', '物联网', 'RFID技术可识别高速运动物体并可同时识别多个标签，操作快捷方便。短距离射频产品不怕油渍、灰尘污染等恶劣的环境，可在这样的环境中替代条码，例如用在工厂的流水线上跟踪物体。长距射频产品多用于交通上，识别距离可达几十米，如自动收费或识别车辆身份等。', '2019-02-27'),
       (7, '云计算与物联网的关键技术', '物联网世界', '云计算实现了通过网络提供可伸缩、廉价的分布式计算能力。它包括3种典型的服务模式：IaaS(基础设施即服务)，PaaS(平台即服务)，SaaS(软件即服务)。IaaS像一些为游戏公司提供研发设备出租的公司，而steam就是PaaS，steam上的游戏就是SaaS。', '2019-02-14')
       ;

create table product(
id integer primary key autoincrement,
title text,
subtitle text,
wcover text,
hcover text
);

insert into product(id, title,subtitle,wcover, hcover) values
(1, 'RFID便携读写器 __title.en__-UR10', '让手机轻松变身专业数据采集设备', 'prod-rfid.png', 'rfid-scanner.png'),
(2, 'LORA网关与节点系统 __title.en__-LG001', '集灵活、高效、性价比与一身的LoRa网关节点设备', '', 'lora-board-high.png'),
(3, 'RFID温度标签 __title.en__-T1901', '满足医药、食品、环境检测等多种使用场景', 'prod-rfid-temp.jpg', 'rfid-tag-long.jpg'),
(4, '区块链硬件钱包', '多币种、多重安全保障硬件钱包', 'middle-left.png', 'wallet-long.png');
(5, '手机智能额温计', '替代传统额温枪的便携式智能体温与物温测量设备', '01.6.png', '01.12-org.png')fma;



create table solution(
id integer primary key autoincrement,
title text,
subtitle text,
wcover text,
hcover text
);

insert into solution(id, title,subtitle,wcover, hcover) values
(1, 'RFID库存管理服务', '将RFID技术结合我们独立开发的云服务，让您轻松管理产品的生产、运输、储存及盘点等各个阶段', '', 'sol-rfid-warehouse.jpeg'),
(2, 'RFID服装管理系统', '针对服装零售供应链各个环节的管理痛点，定制基于物联网技术的样衣管理系统', '', 'rfidapparel.jpg'),
(3, '“物联网+区块链”可追溯防伪系统', '将区块链技术融合于物联网系统', '', 'blockchain-iot.jpg'),
(4, '智能体脂秤解决方案', '蓝牙/WIFI四、八电极智能体脂秤', '', 'scale-long.jpeg'),
(5, 'LoRa物联网解决方案', '基于LoRa的低功耗、远距离、低成本无线通讯解决方案', '', 'IoT.jpg');

