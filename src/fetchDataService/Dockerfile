FROM python

# Create service directory
RUN mkdir /service
# Create a program logs folder in the service directory
RUN mkdir /service/program\ logs

# Copy the 'proto' and 'interceptors' folders into the service directory
COPY src/fetchDataService/proto/ /service/proto
COPY src/fetchDataService/interceptors/ /service/interceptors
COPY certification/ /service/certification
COPY TestData /service/TestData
COPY src/fetchDataService/requirements.txt /service
COPY src/fetchDataService/fetchServer.py /service

COPY /src/fetchDataService/configuration.yaml /service

# Basically performing cd inside the image and setting the work directory
WORKDIR /service

# Update pip to avoid warnings
RUN python3 -m pip install --upgrade pip
# Tells docker to run pip install inside the image which will add
# all the required packages
RUN python3 -m pip install -r requirements.txt
# Run the python protoc command to generate the python files
# RUN python3 -m grpc_tools.protoc -I ../protobufs --python_out=. / --grpc_python_out=

# Tells Docker that we're going to run the microservice on port 50051
EXPOSE 50051
# Tells Docker how to run the microservice
ENTRYPOINT [ "python3", "fetchServer.py" ]