FROM jrei/systemd-centos:7

COPY ApplicationManagement.viper.yaml /etc/ApplicationManagement/ApplicationManagement.viper.yaml
COPY application-management /usr/local/bin/application-management

CMD ["application-management", "run"]
