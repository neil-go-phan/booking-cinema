import axiosClient from '@/helpers/axiosClient';

export async function checkLogin() {
  try {
    const res: any = await axiosClient.get(`/auth/check-login`);
    if (res.data.success) {
      return true;
    }
    return false;
  } catch (error) {
    return false;
  }
}