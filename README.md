# WildScribe

[Team Miro Board](https://miro.com/app/board/uXjVNZAgD-U=/) 

[WildScribe Organization](https://github.com/2305-WildScribe)

## Development Teams

### Front End Team

- Alice Abarca: [GitHub](https://github.com/aliceabarca/) | [LinkedIn](https://www.linkedin.com/in/alice-abarca-431615272/)
- Jocelyn Wensloff: [GitHub](https://github.com/Jwensloff/) | [LinkedIn](https://www.linkedin.com/in/jocelynwensloff/)

### Back End Team

- Parker Boeing: [GitHub](https://github.com/ParkerBoeing) | [LinkedIn](https://www.linkedin.com/in/parker-boeing/)
- Derek Chavez: [GitHub](https://github.com/DChavez18) | [LinkedIn](https://www.linkedin.com/in/derek-chavez/)
- Ian Lyell: [GitHub](https://github.com/ILyell) | [LinkedIn](https://www.linkedin.com/in/ian-lyell/)

---

## Installing Go on macOS

Go (also known as Golang) is a popular programming language developed by Google. This guide will walk you through the steps to install Go on a macOS system.

## Prerequisites

Before you begin, make sure you have the following:

- A macOS computer
- An internet connection

## Installation Steps

Follow these steps to install Go on your macOS system:

1. **Download the Go Installer:**

   Visit the official Golang website to download the installer for macOS. Go to [https://golang.org/dl/](https://golang.org/dl/) and find the macOS version.

2. **Choose the Correct Package:**

   You will see several options for macOS, typically labeled as `goX.Y.darwin-amd64.pkg`, where `X.Y` represents the version number. Click on the latest version to download it.

3. **Install Go:**

   After downloading the package, open the downloaded `.pkg` file by double-clicking it. This will initiate the installation process.

4. **Follow the Installer Instructions:**

   Follow the on-screen instructions to complete the installation. You may need to enter your password to allow the installer to make changes to your system.

5. **Verify the Installation:**

   After the installation is complete, open your terminal. To verify that Go has been installed successfully, you can run the following command:

   ```bash
   go version

---

### How to Install the Project

- Fork and clone this repo
- Run go test in order to run the test suite

---

# API JSON Contract

## Users

Description of API endpoints for Front End application

### Getting User

`POST /api/v0/user`

**Request**

```json
{
    "data": {
        "type": "user",
        "attributes": {
            "email": "me@gmail.com",
            "password": "hi"
        }
    }
}
```
**Success Response (200 OK)**:

- **Status**: 200 OK
- **Description**: Successful response with user id
- **Data format**: a hash with a hash of user data

```json
{
    "data": {
        "type": "user",
        "attributes": {
            "name": "Ian",
            "user_id": "652edaa67a75034ea37c6652"
        }
    }
}
```
## Adventures

Description of API endpoints for Front End application

### Getting Adventures for User

`POST /api/v0/user/adventures`

**Request**

```json
{
    "data":{
        "type": "adventure",
        "attributes":{
            "adventure_id": "652da923ff996de855a6d39d"
        }
    }
}
```
**Success Response (200 OK)**:

- **Status**: 200 OK
- **Description**: Successful response with all adventures associated with user id
- **Data format**: a hash with all adventures, with a hash of adventure data

```json
{
    "data": {
        "type": "adventure",
        "attributes": {
            "user_id": "65299d4ceb708107b33729c6",
            "adventure_id": "652da923ff996de855a6d39d",
            "activity": "Running",
            "date": "10/11/2023",
            "image_url": "https://www.rei.com/dam/parrish_091412_0679_main_lg.jpg",
            "stress_level": "Very High",
            "hours_slept": 8,
            "sleep_stress_notes": "notes about sleep and stress",
            "hydration": "Hydrated",
            "diet": "Good Diet",
            "diet_hydration_notes": "Some Hydraytion",
            "beta_notes": "Running is real hard"
        }
    },
    {
        "type": "adventure",
        "attributes": {
            "user_id": "65299d4ceb708107b33729c6",
            "adventure_id": "652da923ff996de855a6d39d",
            "activity": "Swimming",
            "date": "10/11/2024",
            "image_url": "https://www.rei.com/dam/parrish_091412_0679_main_lg.jpg",
            "stress_level": "High",
            "hours_slept": 9,
            "sleep_stress_notes": "notes about sleep and stress",
            "hydration": "Hydrated",
            "diet": "Good Diet",
            "diet_hydration_notes": "Some Hydraytion",
            "beta_notes": "Swimming is real hard"
        }
    }
}
```
### Getting An Adventure

`POST /api/v0/user/adventure`

**Request**

```json
{
    "data":{
        "type": "adventure",
        "attributes":{
            "adventure_id": "652ff8c82ed41a2d015d993b"
        }
    }
}
```
**Success Response (200 OK)**:

- **Status**: 200 OK
- **Description**: Successful response with adventure data associated with adventure id
- **Data format**: a hash with adventure, with a hash of adventure data

```json
{
    "data": {
        "type": "adventure",
        "attributes": {
            "user_id": "65299d4ceb708107b33729c6",
            "adventure_id": "652ff8c82ed41a2d015d993b",
            "activity": "Running",
            "date": "10/11/2023",
            "image_url": "https://www.rei.com/dam/parrish_091412_0679_main_lg.jpg",
            "stress_level": "Very High",
            "hours_slept": 8,
            "sleep_stress_notes": "notes about sleep and stress",
            "hydration": "Hydrated",
            "diet": "Good Diet",
            "diet_hydration_notes": "Some Hydraytion",
            "beta_notes": "Running is real hard"
        }
    }
}
```
### Creating An Adventure

`POST /api/v0/adventure`

**Request**

```json
{
 "data": {
        "type": "adventure",
        "attributes": {
            "user_id": "65299d4ceb708107b33729c6",
            "activity": "Running",
            "date": "10/11/2023",
            "notes": "Running is hard",
            "image_url": "https://www.rei.com/dam/parrish_091412_0679_main_lg.jpg",
            "stress_level": "Very High",
            "hours_slept": 8,
            "sleep_stress_notes": "notes about sleep and stress",
            "hydration": "Hydrated",
            "diet": "Good Diet",
            "diet_hydration_notes": "Some Hydraytion",
            "beta_notes": "Running is real hard"
        }
    }
}
```
**Success Response (201 OK)**:

- **Status**: 201 OK
- **Description**: Successful response with adventure id and success message
- **Data format**: a hash with message, with a hash of new adventure id

```json
{
    "data": {
        "type": "adventure",
        "message": "success",
        "attributes": {
            "adventure_id": "652ff8c82ed41a2d015d993b"
        }
    }
}
```
### Deleting An Adventure

`DELETE /api/v0/adventure`

**Request**

```json
{
    "data": {
        "type": "adventure",
        "attributes": {
            "adventure_id":"6530428eb4e1886116236a8a"
        }
    }
}
```
**Success Response (200 OK)**:

- **Status**: 200 OK
- **Description**: Successful response with success message
- **Data format**: a hash with message and and adventure type

```json
{
    "data": {
        "type": "adventure",
        "message": "success"
    }
}
```
### Updating An Adventure

`PUT /api/v0/adventure`

**Request**

```json
{
 "data":{
        "type": "adventure",
        "attributes":{ 
            "adventure_id": "12",
            "activity": "Running",
            "date": "10/11/2023",
            "notes": "Running is hard",
            "image_url": "https://www.rei.com/dam/parrish_091412_0679_main_lg.jpg",
            "stress_level": "Very High",
            "hours_slept": 8,
            "sleep_stress_notes": "notes about sleep and stress",
            "hydration": "Hydrated",
            "diet": "Good Diet",
            "diet_hydration_notes": "Some Hydraytion",
            "beta_notes": "Running is real hard"
        }
    }
}
```
