<style scoped>
	.protogo-content {
		display: flex;
		align-items: center;
		justify-content: center;
		height: 100%;
	}

	.full-height {
		height: 100vh;
	}
</style>
<template>
<ion-tabs>
	<ion-tab tab="profile">
		<div id="me-page">
			<ion-header>
				<ion-toolbar>
					<ion-title>Profile</ion-title>
				</ion-toolbar>
			</ion-header>
			<ion-content class="full-height">
				<ion-card>
					<img alt="Silhouette of mountains" src="https://ionicframework.com/docs/img/demos/card-media.png" />
					<ion-card-header>
						<ion-avatar>
							<img alt="Silhouette of a person's head" src="https://ionicframework.com/docs/img/demos/avatar.svg" />
						</ion-avatar>
						<ion-card-title>Rostam Pesar'zaal</ion-card-title>
						<ion-card-subtitle>@rostam</ion-card-subtitle>
					</ion-card-header>
					<ion-card-content>
						Monster hunter.
					</ion-card-content>
				</ion-card>
			</ion-content>
		</div>
	</ion-tab>

	<ion-tab tab="gigs">
		<div id="gigs-page">
			<ion-header>
				<ion-toolbar>
					<ion-title>Gigs</ion-title>
				</ion-toolbar>
			</ion-header>
			<ion-content class="full-height">
				<div class="protogo-content">Gigs content</div>
			</ion-content>
		</div>
	</ion-tab>

	<ion-tab tab="search">
		<div id="search-page">
			<ion-header>
				<ion-toolbar>
					<ion-title>Search</ion-title>
				</ion-toolbar>
			</ion-header>
			<ion-content class="full-height">
				<div class="protogo-content">Search content</div>
			</ion-content>
		</div>
	</ion-tab>

	<ion-tab tab="chat">
		<div id="chats-page">
			<ion-header>
				<ion-toolbar>
					<ion-title>Chat</ion-title>
				</ion-toolbar>
			</ion-header>
			<ion-content class="full-height">
				<div class="protogo-content">Chat content</div>
			</ion-content>
		</div>
	</ion-tab>

	<ion-tab tab="connections">
		<div id="connections-page">
			<ion-header>
				<ion-toolbar>
					<ion-title>Connections</ion-title>
				</ion-toolbar>
			</ion-header>
			<ion-content class="full-height">
				<div v-for="connectionGroup in connectionGroups">
					<ion-item-divider>
						<ion-label>{{connectionGroup.name}}</ion-label>
					</ion-item-divider>

					<div v-for="connection in connectionGroup.members">
						<ion-item :key="connection.id">
							<ion-avatar>
								<ion-icon :icon="personOutline" size="large"></ion-icon>
							</ion-avatar>
							<ion-label>{{ connection.name }}</ion-label>
						</ion-item>
					</div>
				</div>


				<ion-item-group>
					<ion-item-divider>
						<ion-label>FOSDEM 2025</ion-label>
					</ion-item-divider>

					<ion-item>
						<ion-label>A</ion-label>
					</ion-item>
					<ion-item>
						<ion-label>BB</ion-label>
					</ion-item>
					<ion-item>
						<ion-label>CCC</ion-label>
					</ion-item>
				</ion-item-group>


				<!--
				<ion-list class="full-height">
					<div v-for="item in items">
					<ion-item :key="item.id">
						<ion-avatar>
							<ion-icon :icon="personOutline" size="large"></ion-icon>
						</ion-avatar>
						<ion-label>{{ item.name }}</ion-label>
					</ion-item>
					</div>
				</ion-list>
				-->
			</ion-content>
		</div>
	</ion-tab>

	<ion-tab-bar slot="bottom">
		<ion-tab-button tab="profile">
			<ion-icon :icon="personCircleOutline" />
			Profile
		</ion-tab-button>
		<ion-tab-button tab="gigs">
			<ion-icon :icon="briefcaseOutline" />
			Gigs
		</ion-tab-button>
		<ion-tab-button tab="search">
			<ion-icon :icon="searchOutline" />
			Search
		</ion-tab-button>
		<ion-tab-button tab="chat">
			<ion-icon :icon="chatboxOutline" />
			Chat
		</ion-tab-button>
		<ion-tab-button tab="connections">
			<ion-icon :icon="atOutline" />
			Connections
		</ion-tab-button>
	</ion-tab-bar>
</ion-tabs>
</template>

<script lang="ts">
import {
	IonAvatar,
	IonCard,
	IonCardContent,
	IonCardHeader,
	IonCardSubtitle,
	IonCardTitle,
	IonContent,
	IonHeader,
	IonIcon,
	IonItem,
	IonItemDivider,
	IonItemGroup,
	IonLabel,
	IonList,
	IonTab,
	IonTabBar,
	IonTabButton,
	IonTabs,
	IonTitle,
	IonToolbar,
} from '@ionic/vue';

import {
	atOutline,
	briefcaseOutline,
	chatboxOutline,
	personOutline,
	personCircleOutline,
	searchOutline,
} from 'ionicons/icons';

const connectionGroups = [
	{id:"1",name:"Electronic Arts (EA), Current & Alumni",members:[
		{id:"1",name:"Ani Lopez"},
		{id:"2",name:"Byron Jung"},
		{id:"3",name:"Edmond Chan"},
		{id:"4",name:"Edward Lee"},
		{id:"5",name:"Charles Iliya Krempeaux"},
		{id:"6",name:"Chris Bon"},
		{id:"7",name:"Gavin McDonald"},
		{id:"8",name:"Jason Paul Moore"},
		{id:"9",name:"Jose (PePe) Amengual"},
		{id:"10",name:"Kevin Harley"},
		{id:"11",name:"Mathieu Desjarlais"},
		{id:"12",name:"Michael Wu"},
		{id:"13",name:"Mohamed El Eryan"},
		{id:"14",name:"Radoslav Gazo"},
		{id:"15",name:"Rishabh Bhatia"},
		{id:"16",name:"Samar Hashemi"},
		{id:"17",name:"Stephen Carr"},
		{id:"18",name:"Tommy Chan"},
	]},
	{id:"2",name:"Fuerza-T",members:[
		{id:"1",name:"Carnivor"},
		{id:"2",name:"Destructor"},
		{id:"3",name:"Felinor"},
		{id:"4",name:"Goldar"},
		{id:"5",name:"Infector"},
		{id:"6",name:"Pantano"},
		{id:"7",name:"Platino"},
		{id:"8",name:"Shidoshi"},
		{id:"9",name:"Silvar"},
		{id:"10",name:"Spector"},
		{id:"11",name:"Tekno"},
		{id:"12",name:"Titano"},
	]},
	{id:"3",name:"GreatApe",members:[
		{id:"1",name:"Atanaz Ostovar"},
		{id:"2",name:"Benyamin Azarkhazin"},
		{id:"3",name:"Charles Iliya Krempeaux"},
		{id:"4",name:"Chet Earl Woodside"},
		{id:"5",name:"Chris Trottier"},
		{id:"6",name:"Farzaneh Amini"},
		{id:"7",name:"Massoud Seifi"},
		{id:"8",name:"Mehrdad Mirsamie"},
		{id:"9",name:"Meysam Mousavi"},
		{id:"10",name:"Muhammad Zaid Ali"},
		{id:"11",name:"Nariman Movaffaghi"},
		{id:"12",name:"Sal Rahman"},
		{id:"13",name:"Sepideh Farsi"},
	]},
	{id:"4",name:"Mysterians",members:[
		{id:"1",name:"Blazar"},
		{id:"1",name:"Jeepar"},
		{id:"1",name:"Truckar"},
	]},
];

export default {
	components: {
		IonAvatar,
		IonCard,
		IonCardContent,
		IonCardHeader,
		IonCardSubtitle,
		IonCardTitle,
		IonContent,
		IonHeader,
		IonIcon,
		IonItem,
		IonItemDivider,
		IonItemGroup,
		IonLabel,
		IonList,
		IonTab,
		IonTabBar,
		IonTabButton,
		IonTabs,
		IonTitle,
		IonToolbar,
	},
	data() {
		return {
			atOutline,
			briefcaseOutline,
			chatboxOutline,
			searchOutline,
			personCircleOutline,
			personOutline,

			connectionGroups,
		};
	},
};
</script>
