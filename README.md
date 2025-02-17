# Parse-For-MerkleBot

## Overview

**Parse-For-MerkleBot** is a project designed to automate the process of collecting information about robotics companies and their GitHub repositories related to ROS (Robot Operating System). The project includes a set of scripts for data parsing, link searching, and repository analysis.

## Table of Contents
- [Project Description](#project-description)
- [Scripts](#scripts)
  - [searchGoogleAPI.go](#searchgoogleapigo)
  - [filter.go](#filtergo)
  - [GetTop.go](#gettopgo)
- [Data Storage](#data-storage)

## Project Description

### 1. Collecting Company Information
- Extracting a list of companies and their websites from the [README file](https://github.com/vmayoral/ros-robotics-companies) of the `ros-robotics-companies` project.
- The results are saved in the `companies.txt` file.

### 2. Searching for GitHub Links
- Using the Google API to search for all GitHub links related to company names from `companies.txt`.
- The results are recorded in the `Git List.txt` file.

### 3. Filtering Results
- The `filter.go` script processes the `Git List.txt` file, extracting possible links to the companies' GitHub profiles.

### 4. Analyzing Repositories
- The `GetTop.go` script outputs the top repositories containing the word "ROS" in their description to the console.

## Scripts

### searchGoogleAPI.go

This script parses the `companies.txt` file and searches Google for all GitHub links related to the company name. The results are saved in the `Git List.txt` file.

### filter.go

This script parses the `Git List.txt` file and extracts possible links to the company's GitHub profile.

### GetTop.go

This script outputs the top repositories containing the word "ROS" in their description to the console.

## Data Storage

All results are also transferred to a Google Spreadsheet with three sheets for easy viewing and analysis. 

📌 [Google Spreadsheet](https://docs.google.com/spreadsheets/d/1vKJp_UpHgkx5pog60oYFNFI7powGT5rZJlZW7NhUefA/edit?usp=sharing)
