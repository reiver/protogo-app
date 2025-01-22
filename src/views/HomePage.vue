<style scoped>
	.protogo-content {
		display: flex;
		align-items: center;
		justify-content: center;
		height: 100%;
	}
</style>
<template>
<ion-tabs>
	<ion-tab tab="me">
		<div id="me-page">
			<ion-header>
				<ion-toolbar>
					<ion-title>Me</ion-title>
				</ion-toolbar>
			</ion-header>
			<ion-content>
				<div class="protogo-content">My content</div>
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
			<ion-content>
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
			<ion-content>
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
			<ion-content>
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
			<ion-content>
				<div class="protogo-content">Connections content</div>
			</ion-content>
		</div>
	</ion-tab>

	<ion-tab-bar slot="bottom">
		<ion-tab-button tab="me">
			<ion-icon :icon="bodyOutline" />
			Me
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
	IonTabs,
	IonTab,
	IonToolbar,
	IonTabBar,
	IonTabButton,
	IonContent,
	IonHeader,
	IonTitle,
	IonIcon,
} from '@ionic/vue';

import { atOutline, bodyOutline, briefcaseOutline, chatboxOutline, searchOutline, } from 'ionicons/icons';

export default {
	components: { IonTabs, IonTab, IonToolbar, IonTabBar, IonTabButton, IonContent, IonHeader, IonTitle, IonIcon },
	data() {
		return {
			atOutline,
			bodyOutline,
			briefcaseOutline,
			chatboxOutline,
			searchOutline,
		};
	},
};
</script>
