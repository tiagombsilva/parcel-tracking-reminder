# Building and tests
FROM openjdk:21-slim AS build

WORKDIR /app

COPY . /app/

RUN chmod +x ./mvnw

RUN ./mvnw install

# Final image using distroless
FROM gcr.io/distroless/java21-debian12

WORKDIR /app

COPY --from=build /app/target/*.jar app.jar

ENTRYPOINT ["java", "-jar", "/app/app.jar"]