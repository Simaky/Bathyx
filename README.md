[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![MIT License][license-shield]][license-url]



<!-- PROJECT LOGO -->
<br />
<div align="center">
  <a href="https://github.com/Simaky/Bathyx">
    <img src=".github/logo.png" alt="Logo" width="80" height="80">
  </a>

<h3 id="readme-top" align="center">Bathyx</h3>

  <p align="center">
    Open source library and tray application that show battery percentage for supported devices
   <!-- <br />
    <a href="https://github.com/Simaky/Bathyx/issues"><strong>Explore the docs »</strong></a>
 -->    
<br />
    <br />
    <a href="https://github.com/Simaky/Bathyx/issues">Report Bug</a>
    ·
    <a href="https://github.com/Simaky/Bathyx/issues">Request Feature</a>
  </p>
</div>



<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#roadmap">Roadmap</a></li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#license">License</a></li>
  </ol>
</details>



<!-- ABOUT THE PROJECT -->
## About The Project

![Bathyx Screen Shot][product-screenshot]

Open source library and tray application that show battery percentage for supported devices.
For now - I add devices that I have and I am using every day.
If you want me to add your device, please create an request issue and I'll check what could I do.

Implemented features:
* Cross platform tray application with battery percentage
* Open source library that give you ability to use device data in your project, just import this lib
* Support HyperX Cloud Flight S headphones

<p align="right">(<a href="#readme-top">back to top</a>)</p>



### Built With


<img src="https://upload.wikimedia.org/wikipedia/commons/thumb/0/05/Go_Logo_Blue.svg/512px-Go_Logo_Blue.svg.png?20191207190041" width="100">

Special thanks [karalabe/hid](https://github.com/karalabe/hid) ([forked](https://github.com/Simaky/hid-v2)) and [getlantern/systray](github.com/getlantern/systray)

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- GETTING STARTED -->
## Getting Started

### Application

To download application please find the latest **[Release](https://github.com/Simaky/Bathyx/releases)** and choose build for your OS.

### Library

To use Bathyx as library just simply do:

1. Get and Import app
   ```sh
   go get github.com/Simaky/Bathyx
   ```

   ```sh
   import "github.com/Simaky/Bathyx/devices"
   ```
2. Get devices and receive information
   ```sh
   // new devices
   d :=  devices.New()
   
   // receive information
   info := <-d.HyperX.CloudFlightS(ctx, time.Minute*1)
   ```

<p align="right">(<a href="#readme-top">back to top</a>)</p>


<!-- ROADMAP -->
## Roadmap

- [x] Add HyperX Cloud Flight S support
- [x] Add Tray application
- [x] Add README.md
- [x] Add "components" document to easily copy & paste sections of the readme
- [ ] Support more devices


See the [open issues](https://github.com/othneildrew/Best-README-Template/issues) for a full list of proposed features (and known issues).

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- CONTRIBUTING -->
## Contributing

Contributions are what make the open source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

If you have a suggestion that would make this better, please fork the repo and create a pull request. You can also simply open an issue with the tag "enhancement".
Don't forget to give the project a star! Thanks again!

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- LICENSE -->
## License

Distributed under the MIT License. See `LICENSE.txt` for more information.

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[contributors-shield]: https://img.shields.io/github/contributors/Simaky/Bathyx.svg?style=for-the-badge
[contributors-url]: https://github.com/Simaky/Bathyx/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/Simaky/Bathyx.svg?style=for-the-badge
[forks-url]: https://github.com/Simaky/Bathyx/network/members
[stars-shield]: https://img.shields.io/github/stars/Simaky/Bathyx.svg?style=for-the-badge
[stars-url]: https://github.com/Simaky/Bathyx/stargazers
[issues-shield]: https://img.shields.io/github/issues/Simaky/Bathyx.svg?style=for-the-badge
[issues-url]: https://github.com/Simaky/Bathyx/issues
[license-shield]: https://img.shields.io/github/license/Simaky/Bathyx.svg?style=for-the-badge
[license-url]: https://github.com/Simaky/Bathyx/blob/master/LICENSE.txt
[product-screenshot]: .github/screenshot.png