# golang-admin-brute

The project was created to crawl the site in search of the entrance to the administrative panel.
The idea of the project is for testers who are looking for vulnerabilities. The developer is not responsible for your actions.
## Setup
Clone the repository and change the working directory:

    git clone https://github.com/nikitavoryet/golang-admin-brute.git
    cd golang-admin-brute

Build and run the program:

    go build -o brute
    ./brute

# In the plans :
```
- [ ] Add proxy
- [ ] Mass search
- [ ] Use go rutina
- [ ] Api with generate report to tg/mail
```

# How does it work :
```
Get URL with path from file adminpath.txt
If the answer is !== 404 -> this domain with path save to file
The file is generated automatically from the current timetamp and _output.txt
```

```
Donate:

    BTC:  192TC7d7ZRYJQbQnAWvMpkccnBNQN1ae6R
    ETH:  0x7d1082d952f4d584ae2910e14018f4dce7495c74
    LTC:  MLx6wmFjXfBTKj6JfB5NXaiKjNLeEntRoZ
    DOGE: DHCjW71EWBzvv43XPXVJc491brcBJXXq88
```
    author: 
    
    Name:          Nikita
    Company:       SmartWorld
    Position:      TeamLead
    Mail:          n.vtorushin@inbox.ru
    TG:            @nikitavoryet
    Year of birth: 1999
    FullStack:     JS/GO