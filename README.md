<!-- Improved compatibility of remonter link: See: https://github.com/othneildrew/Best-README-Template/pull/73 -->

<a name="readme-top"></a>

<!--
*** Thanks for checking out the Best-README-Template. If you have a suggestion
*** that would make this better, please fork the repo and create a pull request
*** or simply open an issue with the tag "enhancement".
*** Don't forget to give the project a star!
*** Thanks again! Now go create something AMAZING! :D
-->

<!-- PROJECT SHIELDS -->
<!--
*** I'm using markdown "reference style" links for readability.
*** Reference links are enclosed in brackets [ ] instead of parentheses ( ).
*** See the bottom of this document for the declaration of the reference variables
*** for contributors-url, forks-url, etc. This is an optional, concise syntax you may use.
*** https://www.markdownguide.org/basic-syntax/#reference-style-links
-->

<!-- PROJECT LOGO -->
<br />
<div align="center">

<h1 align="center">[PIDR] EnergyWatch</h1>

  <p align="center">
    Supervision énergétique
    <br />
    <a href="https://github.com/PIDR-EnergyWatch/EnergyWatch/-/tree/main/docs"><strong>Explorer les documents »</strong></a>
    <br />
    <br />
    <a href="https://github.com/PIDR-EnergyWatch/EnergyWatch"><img alt="pipeline status" src="https://github.com/PIDR-EnergyWatch/EnergyWatch/actions/workflows/docker-image.yml/badge.svg" /></a>
    <br />
    <a href="https://github.com/PIDR-EnergyWatch/EnergyWatch/issues">Signaler un Bug</a>
    ·
    <a href="https://github.com/PIDR-EnergyWatch/EnergyWatch/issues">Demander une fonctionnalité</a>
  </p>
</div>

<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table des matières</summary>
  <ol>
    <li>
      <a href="#about-the-project">À propos du projet</a>
      <ul>
        <li><a href="#built-with">Fait avec</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Démarrer</a>
      <ul>
        <li><a href="#prerequisites">Prérequis</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#usage">Utilisation</a></li>
    <li><a href="#roadmap">Roadmap</a></li>
    <li><a href="#contributing">Contributions</a></li>
    <li><a href="#license">Licence</a></li>
    <li><a href="#contact">Contact</a></li>
    <li><a href="#acknowledgments">Remerciements</a></li>
  </ol>
</details>

<!-- ABOUT THE PROJECT -->

## À propos du projet

Ce repo présente un projet de collaboration entre la société de génie électrique CEGELEC Lorraine, l’ENSEM et TELECOM Nancy, visant à développer un système de supervision de la consommation énergétique pour les bâtiments universitaires. L’objectif principal du projet est de sensibiliser les personnels et les étudiants à la transition énergétique et aux écogestes en affichant en temps réel les consommations électriques, la température extérieure, l’hygrométrie, la consommation due au chauffage et l’impact carbone des différentes activités dans le hall de l’école. Le système de supervision repose sur l’utilisation de capteurs sans fil pour mesurer les données, ainsi que sur une interface utilisateur conviviale pour visualiser et analyser les informations. La solution a été conc ̧ue pour être facilement dupliquée dans les bâtiments d’autres écoles et/ou d’autres secteurs d’activité

<p align="right">(<a href="#readme-top">remonter</a>)</p>

### Fait avec

- [![Docker](https://img.shields.io/badge/Docker-2496ED?logo=docker&logoColor=fff)](#)
- [![GitHub Actions](https://img.shields.io/badge/GitHub_Actions-2088FF?logo=github-actions&logoColor=white)](#)
- [![Svelte](https://img.shields.io/badge/Svelte-%23f1413d.svg?logo=svelte&logoColor=white)](#)
- [![Go](https://img.shields.io/badge/Go-%2300ADD8.svg?&logo=go&logoColor=white)](#)
- [![InfluxDB](https://img.shields.io/badge/InfluxDB-22ADF6?logo=influxdb&logoColor=fff)](#)

<p align="right">(<a href="#readme-top">remonter</a>)</p>

<!-- GETTING STARTED -->

## Démarrer

Pour mettre en place une copie locale et la faire fonctionner, suivez ces étapes simples.

### Prérequis

1. Docker & docker-compose

### Installation

1. Cloner le repo
   ```sh
   git clone https://github.com/PIDR-EnergyWatch/EnergyWatch.git
   ```
2. Compiler
   ```sh
   cd mezureux1u
   make build
   ```
3. Lancer
   ```sh
   make up
   ```

<p align="right">(<a href="#readme-top">remonter</a>)</p>

<!-- USAGE EXAMPLES -->

## Utilisation

### Interface graphique

Rendez-vous à l'adresse <a href="http://localhost:3000/">http://localhost:3000/</a>

### Adresses des services intermédiaires

1. Base de données : <a href="http://localhost:8086/">http://localhost:8086/</a>

2. API : <a href="http://localhost:8000/">http://localhost:8000/</a>

## Roadmap

- [x] Persistance des données

  - [x] MQTT

  - [x] API externe

  - [ ] de l'UL

- [x] Écran de supervision

<p align="right">(<a href="#readme-top">remonter</a>)</p>

<!-- CONTRIBUTING -->

## Contributions

Les contributions sont ce qui fait de la communauté open source un endroit extraordinaire pour apprendre, inspirer et créer. Toutes vos contributions sont **très appréciées**.

Si vous avez une suggestion qui permettrait d'améliorer ce projet, merci de forker le repo et de créer une pull request. Vous pouvez aussi simplement ouvrir un problème avec le tag « enhancement ».
N'oubliez pas de donner une étoile au projet ! Merci à tous !

1. Forker le projet
2. Créer une branche pour la fonctionnalité (`git checkout -b feature/AmazingFeature`)
3. Commit ses fonctionnalités (`git commit -m 'Add some AmazingFeature'`)
4. Push sur la branche (`git push origin feature/AmazingFeature`)
5. Ouvrir une Pull Request

<p align="right">(<a href="#readme-top">remonter</a>)</p>

<!-- LICENSE -->

## Licence

Distribué sous la licence MIT. Voir `LICENSE.txt` pour plus d'informations.

<p align="right">(<a href="#readme-top">remonter</a>)</p>

<!-- CONTACT -->

## Contact

- Adrien LAROUSSE <<adrien.larousse@telecomnancy.eu>>
- Stanislas MEZUREUX <<stanislas.mezureux@telecomnancy.eu>>

Lien du projet : [https://github.com/PIDR-EnergyWatch/EnergyWatch](https://github.com/PIDR-EnergyWatch/EnergyWatch)

<p align="right">(<a href="#readme-top">remonter</a>)</p>

<!-- ACKNOWLEDGMENTS -->

## Remerciements

- [Chart.js](https://www.chartjs.org)

<p align="right">(<a href="#readme-top">remonter</a>)</p>

<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
