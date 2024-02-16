import { makeRequest } from "$lib/api";
export async function createGroup(groupData,cookies) {
    try {
        const response = await makeRequest('createGroup', 'post', groupData, {},cookies);

        if (response.status === 201) {
            return response.data;
        } else {
            console.error('Failed to create group:', response.data);
            throw new Error(response.data.message || 'Failed to create group');
        }
    } catch (error) {
        console.error('Error creating group:', error);
        throw error;
    }
}