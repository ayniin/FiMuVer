const API_BASE_URL = import.meta.env.VITE_API_BASE_URL || 'http://localhost:8080/api/v1';

const getAuthHeaders = () => {
    const user = JSON.parse(localStorage.getItem('user') || '{}');
    return {
        'Content-Type': 'application/json',
        ...(user.token ? { Authorization: `Bearer ${user.token}` } : {}),
    };
};

const InviteAPI = {
    async generateCode(maxUses = 0, expiresAt = null) {
        const body = { max_uses: maxUses };
        if (expiresAt) body.expires_at = expiresAt;

        const res = await fetch(`${API_BASE_URL}/invite/generate`, {
            method: 'POST',
            headers: getAuthHeaders(),
            body: JSON.stringify(body),
        });
        const data = await res.json();
        if (!res.ok) throw new Error(data.error || `HTTP ${res.status}`);
        return data.data;
    },

    async listCodes() {
        const res = await fetch(`${API_BASE_URL}/invite/list`, {
            headers: getAuthHeaders(),
        });
        const data = await res.json();
        if (!res.ok) throw new Error(data.error || `HTTP ${res.status}`);
        return data.data;
    },

    async deleteCode(id) {
        const res = await fetch(`${API_BASE_URL}/invite/${id}`, {
            method: 'DELETE',
            headers: getAuthHeaders(),
        });
        const data = await res.json();
        if (!res.ok) throw new Error(data.error || `HTTP ${res.status}`);
        return data;
    },
};

export default InviteAPI;
