import { AppUser } from "./auth/provider";

export default function ProfilePage({ user }: { user: AppUser }) {
  return <div>Welcome, {user.alias}</div>;
}
