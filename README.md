# Railgadi Management System

## Description

This is a REST API made using GoFr framework.

## Features

- Add features

## API-Documentation

Below are the endpoints available in the API (all of them have hyperlink to jump to their description):

- [`GET /getAllTrains`](#1-get-getalltrains): Get List of all trains.
- [`GET /getTrainByNumber/{n}`](#2-get-gettrainbynumbern): Get train details by train number.
- [`POST /addNewTrain`](#3-post-addnewtrain): Add a new train.
- [`PUT /updateTrainByNumber/{n}`](#4-put-updatetrainbynumbern): Update details of a train.
- [`POST /trainArrival`](#5-post-trainarrival): Make a train arrive at a platform.
- [`PUT /trainDeparture`](#6-put-traindeparture): Make a train depart from a platform.
- [`GET /findTrainOnStationByTrainNumber/{n}`](#7-get-findtrainonstationbytrainnumbern): Locate the train's platform number(if it is on station).
- [`DELETE /deleteTrainByNumber/{n}`](#8-delete-deletetrainbynumbern): Delete a train.
- [`GET /getAllPlatformDetails`](#9-get-getallplatformdetails): Get all Platforms and the name of train standing on them.
- [`POST /createNPlatforms/{n}`](#10-post-createnplatformsn): Create n number of platforms.
- [`DELETE /deleteAllPlatforms`](#11-delete-deleteallplatforms): Delete all platforms.

#### 1. `GET /getAllTrains`

- **Description:**

  - returns list of all trains

- **Response Body:**

![Screenshot from 2023-12-17 02-12-44](https://github.com/adityaslab/zopsmart-task/assets/71747119/ed0d335a-60f8-421b-94c6-29b8eb0f69e2)



#### 2. `GET /getTrainByNumber/{n}`
- **Description:**

  - Get train details by train number.

- **Response Body:**\
  ![Screenshot from 2023-12-17 02-24-59](https://github.com/adityaslab/zopsmart-task/assets/71747119/df4598f2-8739-44fb-ac06-0bd7c2b7daaf)



#### 3. `POST /addNewTrain`

- **Description:**

  - Add a new train.

- **Request Body:**\
  ![Screenshot from 2023-12-17 02-15-19](https://github.com/adityaslab/zopsmart-task/assets/71747119/3cbe18e1-b28a-4076-b921-4e835dc2864e)


- **Response:**\
  ![image](https://github.com/adityaslab/zopsmart-task/assets/71747119/cc282468-6c52-459d-a87f-a124f54ca492)


#### 4. `PUT /updateTrainByNumber/{n}`

- **Description:**
  - Update details of a train.
- **Request Body:**\
  ![Screenshot from 2023-12-17 02-28-07](https://github.com/adityaslab/zopsmart-task/assets/71747119/f74690db-879f-4478-9503-001bd362edd6)

    
- **Response:**
  ![image](https://github.com/adityaslab/zopsmart-task/assets/71747119/5ebbee37-db54-4d97-8d28-84a1fdc495d6)


#### 5. `POST /trainArrival`

- **Description:**

  - Make a train arrive at a platform.

- **Request Body:**\
  ![image](https://github.com/adityaslab/zopsmart-task/assets/71747119/6e1e4a3b-1f22-45b7-943d-7618bfd06255)

- **Response:**
  ![image](https://github.com/adityaslab/zopsmart-task/assets/71747119/28040518-95ec-4c5b-aa3d-82f026d67f3e)


#### 6. `PUT /trainDeparture`

- **Description:**

  - Make a train depart from a platform.

- **Request Body:**\
  ![image](https://github.com/adityaslab/zopsmart-task/assets/71747119/1f7db8e5-80e3-4193-9086-a5dc4cab719a)

- **Response:**
  ![image](https://github.com/adityaslab/zopsmart-task/assets/71747119/4634e593-2587-4f37-b82a-6a325f5fe22e)


#### 7. `GET /findTrainOnStationByTrainNumber/{n}`

- **Description:**

  - Locate the train's platform number(if it is on station).

- **Response:**
  ![image](https://github.com/adityaslab/zopsmart-task/assets/71747119/415e3610-193e-418b-85ef-eea1e328b8ba)


#### 8. `DELETE /deleteTrainByNumber/{n}`

- **Description:**

  - Delete a train.

- **Response:**
  ![image](https://github.com/adityaslab/zopsmart-task/assets/71747119/296cc636-54f4-41af-bc93-2790dfde6cea)


#### 9. `GET /getAllPlatformDetails`

- **Description:**

  - Get all Platforms and the name of train standing on them.

- **Response:**
  ![image](https://github.com/adityaslab/zopsmart-task/assets/71747119/97884ed7-cc34-4be1-bf61-3a291c6959a7)


#### 10. `POST /createNPlatforms/{n}`

- **Description:**

  - Create n number of platforms.

- **Response:**
  ![image](https://github.com/adityaslab/zopsmart-task/assets/71747119/0002afc7-1cc1-4c33-b990-d353b50e424f)


#### 11. `DELETE /deleteAllPlatforms`

- **Description:**

  - Delete all platforms.

- **Response:**
  ![image](https://github.com/adityaslab/zopsmart-task/assets/71747119/fd929b0b-7ded-4bba-b135-039d44e22033)

