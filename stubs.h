#ifdef __cplusplus
extern "C" {
#endif

    #include <stdlib.h>
    #include <stdint.h>

    typedef void * VW_HANDLE;
    typedef void * VW_EXAMPLE;

    VW_HANDLE VW_InitializeA(const char * pstrArgs);

    VW_EXAMPLE VW_ReadExampleA(VW_HANDLE handle, const char * line);
    VW_EXAMPLE VW_EmptyExample(VW_HANDLE handle);
    void VW_ExamplePushFeature(VW_EXAMPLE e, unsigned char ns, uint32_t fid, float v);
    void VW_FinishExample(VW_HANDLE handle, VW_EXAMPLE e);

    //float VW_Learn(VW_HANDLE, void* e);
    float VW_Learn(VW_HANDLE, VW_EXAMPLE e);

    float VW_Get_Weight(VW_HANDLE handle, size_t index, size_t offset);

    float VW_GetPrediction(VW_EXAMPLE e);

    float VW_GetSimpleLabel(VW_EXAMPLE e);
    float VW_GetSimpleLabelPrediction(VW_EXAMPLE e);

    float VW_GetImportance(VW_EXAMPLE e);

    size_t VW_HashFeatureA(VW_HANDLE handle, const char * s, unsigned long u);

    size_t VW_GetNumFeatures(VW_EXAMPLE e);
    float VW_GetLoss(VW_EXAMPLE e);
    //float VW_GetUpdatedPrediction(VW_EXAMPLE e);
    float VW_GetPartialPrediction(VW_EXAMPLE e);
    //float VW_GetTopicPrediction(VW_EXAMPLE e, size_t i);
    double VW_GetSumLoss(VW_HANDLE vw);

    void VW_Finish(VW_HANDLE handle);

#ifdef __cplusplus
}
#endif
