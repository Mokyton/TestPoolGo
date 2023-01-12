FROM	debian:buster

COPY	elastic/ /
COPY    start_server_inside_docker.sh /home

RUN	apt update &&				\
	apt install -y  lsb-base procps curl
RUN	dpkg -i elasticsearch-7.9.2-amd64.deb
RUN	update-rc.d elasticsearch defaults 95 10
RUN	echo "network.host: 0.0.0.0" >> /etc/elasticsearch/elasticsearch.yml &&	\
	echo "http.port: 9200" >> /etc/elasticsearch/elasticsearch.yml &&	\
	echo "discovery.type: single-node" >> /etc/elasticsearch/elasticsearch.yml

ENTRYPOINT ["sh", "./home/start_server_inside_docker.sh"]


