# Copyright 2020 Lingfei Kong <colin404@foxmail.com>. All rights reserved.    
# Use of this source code is governed by a MIT style    
# license that can be found in the LICENSE file.

FROM BASE_IMAGE
MAINTAINER Lingfei Kong <colin404@foxmail.com>

WORKDIR /opt/iam

RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
      echo "Asia/Shanghai" > /etc/timezone && \
      mkdir -p /var/log/iam

COPY iam-authz-server /opt/iam/bin/

ENTRYPOINT ["/opt/iam/bin/iam-authz-server"]
CMD ["-c", "/etc/iam/iam-authz-server.yaml"]
