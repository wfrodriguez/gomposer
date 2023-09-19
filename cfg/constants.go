package cfg

const Logo = `
╔═╗┌─┐┌┬┐┌─┐┌─┐┌─┐┌─┐┬─┐
║═╗│ ││││├─┘│ │└─┐├┤ ├┬┘
╚═╝└─┘┴ ┴┴  └─┘└─┘└─┘┴└─`

const SQLCreateTable = `
create table post(id integer, title text not null, slug text not null, desc text not null, date integer not null, primary key(id autoincrement));
create table tag(tag text, primary key(tag));
create table post_tag(post integer not null, tag text not null,
  constraint tag foreign key(tag) references tag(tag), constraint post foreign key(post) references post(id));`
