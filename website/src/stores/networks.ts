import { ref } from "vue";
import { defineStore } from "pinia";
import { useEnvironmentStore } from "./environment";

export const useNetworksStore = defineStore("networks", {
  state: () => {
    let networks = ref([]);
    
    const { selectedEnvironmentId } = useEnvironmentStore()
    fetch(`/api/v1/environments/${selectedEnvironmentId}/networks`)
      .then(async response => {
        const data = await response.json();
        networks.value = data;
      });

    return { networks: networks }
  },

  getters: {
    count: (state) => state.networks.length,
  },
});
