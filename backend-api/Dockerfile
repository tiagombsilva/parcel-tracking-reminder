FROM openjdk:21-slim as build

WORKDIR /app

COPY target/parcel-tracking-api-0.0.1-SNAPSHOT.jar parcel-tracking-api.jar

EXPOSE 8080

ENTRYPOINT ["java","-jar","parcel-tracking-api.jar"]