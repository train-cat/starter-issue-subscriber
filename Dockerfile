FROM eraac/golang

ADD bin_app /bin_app

CMD ["/bin_app", "-config", "/config.json"]

