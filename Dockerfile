FROM ubuntu:20.04

ARG target
ADD ./build/linux-amd64/${target} /${target}

ARG target
ENV target_bin=$target
CMD ["sh", "-c", "/${target_bin}"]
