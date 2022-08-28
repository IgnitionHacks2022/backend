# Backend

# how to go into docker container
```sh
docker exec -it postgres psql -U postgres
```

## Inspiration
Eight million tonnes of plastic waste leaks into the ocean every year. Scientists predict that by 2040, the volume of waste entering the environment will increase by three times! In fact, only 20 companies contributed more than half of the globe's single-use plastic waste (130 million metric tons) in 2019. Imagine how much it would be now. 

InDaBin is a solution for homeowners and companies to reduce the amount of recyclable material that are sent to landfills. We all know of times when you get busy and don't have the time to properly recycle.


## What it does

InDaBin is a smart garbage sorting system that figures out where your trash should go. First, hold your disposable in front of the camera so it can identify which bin the item should go in. Then, simply drop your item into the chute which will automatically place it in the correct bin!

InDaBin is also equipped with a trash management system. For companies or homeowners who want to track their eco footprint, simply sign up through InDaBin's app that is available on web, Android, and Apple.

On the app, you will be able to see statistics and day to day insights on your trash and recycling habits. Thanks to our Bluetooth detection technology, artificial intelligence and machine learning algorithms, InDaBin can log the type and name of the disposables you throw out.

# Cracking the Judging Criteria

## Technologies

### Web Development
We used Flutter to implement web development. Although our project has many different facets, ranging from hardware to mobile apps, we wanted to include web without forgoing the rest of the technologies.
The most significant technical problem was making sure that getting Bluetooth info worked on all platforms, whether it was Android, Apple, or web (laptops and computers).

### Machine Learning and AI

We used a combination of AI and machine learning APIs, as well as implementing our own algorithms.

We used Google's Vision API and it's label detection and object localization feature to implement detection of the object the person was throwing out. A significant challenge was implementing algorithms to sort the object that was identified to the proper bin.

We also used Google's Text to Speech API to deliver verbal messages to the thrower. A challenge was figuring out what message to deliver (known person registered to the app, unknown person, type of object thrown out), and how to deliver a voice file.

Lastly, we used OpenCV to detect and sense movement in the camera when the person is throwing out trash, so that the backend is only getting useful data. It was challenging because you don't want to call the backend when a shadow shows up on the camera.

### Bluetooth/Hardware

We used Raspberry Pi, a steppe
Although hardware could've been done in Python since it would be easier, we decided to challenge ourselves by writing it in Rust. We implemented a lot of the GPIO and Bluetooth features manually since it wasn't very well documented.

## Design 
 Throughout the building of InDaBin, _Design_ was always kept in mind especially regarding the three factors of user experience, intuitiveness, and ease of use into your project. Three key components from our application that go above and beyond to incorporate these factors are:
 
 1. **Flutter:** Flutter is a cross-platform front-end framework used to give the application ease of use along with great user experience. Firstly, as Flutter compiles natively to Android, iOS, and PC devices, InDaBin is able to run on almost all large consumer based devices giving it great ease of access. Also, since Flutter compiles to the native language of the device it is run on, the user interfaces can be customized to those specific devices increasing the positive impact the application has on user experience.
 2. **Hardware:** For hardware, we used a Raspberry Pi 4 microprocessor to run a DC motor to allow users to drop their waste into one chute for sorting purposes. This leads to a great sense of intuitiveness as the user is not guessing at which bin he has to place the disposable in but instead placing the item smoothly into the single opening as if it were a normal garbage bin. Furthermore, the device has audio feedback to let the user know which bin the item has been placed in. This helps also establish a sense of Environmental awareness. Finally, the entire hardware setup is painted an environmental theme that matches the colour palette of the mobile application further increasing the positive impact on user experience.
 3. **Bluetooth:** Bluetooth is used to give both an intuitive feel while also contributing to ease of use. As the application only cares for devices close to the trash bin (the one throwing out the disposable), bluetooth allows for a more stable connection even in non-networked regions. Furthermore, as the application is just looking for available bluetooth devices, no extra instructions need to be carried out by the user in advance.

 ## Theme
 As spoken in the inspiration above, our project tries to solve one of the large environmental problems present in the world today: **Waste Disposal**. Our project aligns with this category through this problem by creating an apparatus that reduces the environmental impact users of the product have by properly separating out their disposables into correctly recycled bins. Furthermore, through the use of analytics and vocal feedback, the project spreads awareness among its users which further advertises itself as a great solution to this problem. As this project brings awareness to one of the great global environmental issues, it is applicable to the environmental industries.

 ## Originality
 Our product builds upon the well-known idea of garbage sorter, but takes it to a new level of uniqueness through the introduction of vision AI and a statistics management system. This creates a platform of continuous progression with the great benefit of reducing environmental impact. By using image recognition AI, we are able to greatly extract all manual practices in garbage sorting making the process comparable to simple garbage disposal. Furthermore, through the use of statistics and tracking, we are able to introduce the statistics of progression and how much a user is actually contributing to recycling by using our application. This further spreads awareness and also gives self-appreciation to the user for knowing their great impact on the environment. Finally, we also have a vocal feedback assistant that make using the actual product fun and entertaining. This further deviates our product from the norm. making it much more unique than existing solutions for garbage disposal.

## Business Practicality

There are three main target markets for this product:
### Homeowners

Homeowners would want this product the same way people would want fitness tracking apps - to be more aware of their habits, and the impact they are making. Having InDaBin would allow for homeowners to be more aware of recycling, and the impact of the trash they are throwing out. (Maybe Dad is throwing out a lot of beer bottles!)

### Companies

Companies who want to more green would want this product. Similar to how lights are turned off when no one is in office, InDaBin would allow for companies to make eco-friendly decisions based on their stats. For example, replacing plastic straws with paper straws would reduce the amount of items going into the garbage bin.
Companies can also use this as a means of friendly competition, where certain offices/teams can compete with one another to see who has the smaller trash footprint.

### Government

Recently, Toronto banned the use of plastic shopping bags. However, has this really made an impact on plastic bag usage? With InDaBin, municipal, provincial, and even federal governments have the ability to finally put numbers onto trash management and recycling, to make sound decsions on the future of eco friendly procedures and laws.

