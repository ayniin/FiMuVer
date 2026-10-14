import { apiGet, apiPost, apiDelete } from './apiClient';

const InviteAPI = {
    async generateCode(maxUses = 0, expiresAt = null) {
        const body = { max_uses: maxUses };
        if (expiresAt) body.expires_at = expiresAt;

        const response = await apiPost('/invite/generate', body);
        return response?.data;
    },

    async listCodes() {
        const response = await apiGet('/invite/list');
        return response?.data;
    },

    async deleteCode(id) {
        return apiDelete(`/invite/${id}`);
    },
};

export default InviteAPI;
