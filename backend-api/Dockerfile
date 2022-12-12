FROM openjdk:19-jdk-slim as build

WORKDIR /app

COPY mvnw .
COPY .mvn .mvn
COPY pom.xml .

RUN chmod +x ./mvnw
RUN ./mvnw dependency:go-offline -B

COPY src src

RUN ./mvnw package -Dskiptests
RUN mkdir -p target/dependency && (cd target/dependency; jar -xf ../*.jar)
RUN ls target/dependency/META-INF/

#Prd
FROM openjdk:19-jdk-slim as prd
ARG DEPENDENCY=/app/target/dependency

COPY --from=build ${DEPENDENCY}/BOOT-INF/lib /app/lib
COPY --from=build ${DEPENDENCY}/META-INF /app/META-INF
COPY --from=build ${DEPENDENCY}/BOOT-INF/classes /app

EXPOSE 8080

ENTRYPOINT ["java", "-cp", "app:app/lib/*","com.eagle.backendApi.BackendApiApplication"]