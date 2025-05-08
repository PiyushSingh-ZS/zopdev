import { fetchData, postData } from '@/services/api.js';
import { useMutation, useQuery } from '@tanstack/react-query';

export function useGetCloudAccounts(reqParams, options = {}) {
  const temp = {
    data: [
      {
        id: '1',
        name: 'Zop Cloud',
        status: 'READY',
        icon: 'cloud',
        provider: 'gcp',
        lastUpdatedBy: 'owner@zop.dev',
        updatedAt: '28th January 2025, 15:38',
        auditData: {
          all: {
            danger: 15,
            warning: 16,
            pending: 7,
            compliant: 85,
            unchecked: 7,
            total: 130,
          },
          stale: {
            danger: 3,
            warning: 5,
            pending: 2,
            compliant: 18,
            unchecked: 1,
            total: 29,
          },
          overprovision: {
            danger: 2,
            warning: 4,
            pending: 1,
            compliant: 22,
            unchecked: 3,
            total: 32,
          },
          security: {
            danger: 5,
            warning: 3,
            pending: 0,
            compliant: 15,
            unchecked: 2,
            total: 25,
          },
          network: {
            danger: 1,
            warning: 2,
            pending: 1,
            compliant: 10,
            unchecked: 0,
            total: 14,
          },
          storage: {
            danger: 4,
            warning: 2,
            pending: 3,
            compliant: 20,
            unchecked: 1,
            total: 30,
          },
        },
        categoryIcons: {
          stale: 'server',
          overprovision: 'exclamation',
          security: 'shield',
        },
      },
    ],
  };

  return useQuery({
    queryKey: ['resourceAuditGetData', reqParams],
    queryFn: async () => {
      const url = `/cloud-accounts`;
      const data = await fetchData(url, options);
      if (data?.data?.length > 0) return temp;
      return data;
    },
    staleTime: 0,
    cacheTime: 0,
    refetchOnWindowFocus: false,
    refetchOnReconnect: false,
    retry: false,
    ...options,
  });
}

export function usePostAuditData() {
  return useMutation({
    mutationFn: async (req) => {
      let id = req?.id;
      let createResponse = null;

      // Step 1: Conditionally create cloud account
      if (req?.transformedData) {
        const getCloudAccountRes = await postData('/cloud-accounts', req.transformedData);
        id = getCloudAccountRes?.data?.id;

        if (!id) {
          throw new Error('Missing ID in cloud account creation response');
        }

        createResponse = getCloudAccountRes;
      }

      // Step 2: Conditionally trigger audit
      let auditResponse = null;
      if (id && req?.selectedOption) {
        let auditUrl;
        let auditPayload;
        if (req?.selectedOption === 'run-all' || req?.selectedOption === 'all') {
          auditUrl = `/audit/cloud-accounts/${id}/all`;
          auditPayload = {};
        } else {
          auditUrl = `/audit/cloud-accounts/${id}/category/${req?.selectedOption}`;
          auditPayload = req?.selectedOption;
        }

        auditResponse = await postData(auditUrl, auditPayload);
      }

      return {
        createResponse,
        auditResponse,
      };
    },
  });
}
