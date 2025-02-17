import { login, loginWithOAuth, signup } from "@/app/actions";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { Button } from "@/components/ui/button";

export default async function LoginPage() {
  return (
    <form className="flex flex-col gap-3">
      <Label htmlFor="email">Email</Label>
      <Input id="email" name="email" type="email" required />
      <Label htmlFor="password">Password</Label>
      <Input id="password" name="password" type="password" required />
      <div className="flex flex-col gap-3">
        <Button formAction={login}>Login</Button>
        <Button formAction={signup}>Sign up</Button>
        <Button onClick={loginWithOAuth}>Social Login</Button>
      </div>
    </form>
  );
}
