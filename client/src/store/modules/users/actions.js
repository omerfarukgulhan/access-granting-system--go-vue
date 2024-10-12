export default {
  async loadUsers(context) {
    try {
      const response = await fetch(`http://localhost:8080/users`);
      const responseData = await response.json();

      if (!response.ok) {
        throw new Error(responseData.message || "Failed to fetch!");
      }

      const users = responseData.data.map((user) => ({
        id: user.id,
        username: user.username,
        email: user.email,
        profileImage: user.profileImage,
        roles: user.roles.map((role) => ({
          id: role.id,
          name: role.name,
        })),
      }));

      context.commit("setUsers", users);
    } catch (error) {
      console.error("Error fetching users:", error);
      throw error;
    }
  },
};
