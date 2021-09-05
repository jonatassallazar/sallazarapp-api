
-- Criação da tabela de usuários
CREATE TABLE users(
  id int auto_increment primary key,
  name varchar(100) NOT NULL,
  email varchar(100) NOT NULL unique,
  photourl varchar(200),
  accesslevel varchar(30) NOT NULL,
  password varchar(100) NOT NULL,
  updated_at timestamp default current_timestamp(),
  created_at timestamp default current_timestamp()
) ENGINE=INNODB;

-- Criação da tabela de clientes
CREATE TABLE clients(
  id int auto_increment primary key,
  name varchar(100) NOT NULL,
  status varchar(30) NOT NULL,
  email varchar(100),
  phone varchar(10),
  gender varchar(15),
  birthday timestamp,
  updated_at timestamp default current_timestamp(),
  created_at timestamp default current_timestamp()
) ENGINE=INNODB;

-- Criação da tabela de address
CREATE TABLE address(
  id int auto_increment primary key,
  cep int,
  number varchar(10),
  complement varchar(15),
  phone varchar(10),
  neighbourhood varchar(25),
  city varchar(20),
  state varchar(2),
  client_id int NOT NULL,
  FOREIGN KEY (client_id) REFERENCES clients(id)
  ON DELETE CASCADE,
  updated_at timestamp default current_timestamp(),
  created_at timestamp default current_timestamp()
) ENGINE=INNODB;

-- Criação da tabela de products
CREATE TABLE products(
  id int auto_increment primary key,
  name varchar(100) NOT NULL,
  status varchar(30) NOT NULL,
  unity varchar(10),
  weight float,
  photourl varchar(200),
  supplier varchar(100),
  cost int(20) NOT NULL,
  price int(20) NOT NULL,
  updated_at timestamp default current_timestamp(),
  created_at timestamp default current_timestamp()
) ENGINE=INNODB;

-- Criação da tabela de vendas
CREATE TABLE sales(
  id int auto_increment primary key,
  status varchar(30) NOT NULL,
  date timestamp default current_timestamp(),
  shipping int(20),
  discount int(20),
  tax int(20),
  subtotal int(20),
  total int(20),
  payment_method varchar(15),
  installments_quantity int(2),
  observation varchar(500),
  client_id int NOT NULL,
  FOREIGN KEY (client_id) REFERENCES clients(id)
  ON DELETE CASCADE,
  updated_at timestamp default current_timestamp(),
  created_at timestamp default current_timestamp()
) ENGINE=INNODB;

-- Criação da tabela de itens vendidos
CREATE TABLE sold_itens(
  id int auto_increment primary key,
  quantity float,
  total_value int(20),
  product_id int NOT NULL,
  FOREIGN KEY (product_id) REFERENCES products(id)
  ON DELETE CASCADE,
  updated_at timestamp default current_timestamp(),
  created_at timestamp default current_timestamp()
) ENGINE=INNODB;

-- Criação da tabela de parcelas
CREATE TABLE installments(
  id int auto_increment primary key,
  number int(2),
  value int(20),
  due_date timestamp,
  sale_id int NOT NULL,
  FOREIGN KEY (sale_id) REFERENCES sales(id)
  ON DELETE CASCADE,
  updated_at timestamp default current_timestamp(),
  created_at timestamp default current_timestamp()
) ENGINE=INNODB;