FROM openjdk:14-alpine AS app

WORKDIR /app

RUN apk add --no-cache wget
RUN wget https://github.com/Nan1t/NanoLimbo/releases/download/v1.8/NanoLimbo-1.8-all.jar -O NanoLimbo.jar
COPY ./settings.yml .

CMD ["java", "-jar", "NanoLimbo.jar"]