/**
 * 版本相关 API
 */

import { getDemoManagerLatestRelease } from '@/features/demo/demoFixtures';
import { isDemoMode } from '@/features/demo/demoMode';
import { apiClient } from './client';

export interface ManagerLatestRelease {
  tag_name?: string;
  name?: string;
  html_url?: string;
  published_at?: string;
  [key: string]: unknown;
}

export const versionApi = {
  checkLatest: () => apiClient.get<Record<string, unknown>>('/latest-version'),

  checkManagerLatest: async () => {
    if (__DEMO_SITE__ && isDemoMode()) {
      return getDemoManagerLatestRelease();
    }

    return apiClient.get<ManagerLatestRelease>('/latest-manager-version');
  }
};
