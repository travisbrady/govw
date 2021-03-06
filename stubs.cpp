#include <stdio.h>
#include <string>
#include <vowpalwabbit/parser.h>
#include <vowpalwabbit/vw.h>
#include "stubs.h"

extern "C" {

    VW_HANDLE VW_InitializeA(const char * pstrArgs) {
        string s(pstrArgs);
        vw *all = VW::initialize(s);
        return static_cast<VW_HANDLE>(all);
    }

    VW_EXAMPLE VW_ReadExampleA(VW_HANDLE handle, const char * line)
    {
        vw * pointer = static_cast<vw*>(handle);
        // BUGBUG: I really dislike this const_cast. should VW really change the input string?
        return static_cast<VW_EXAMPLE>(VW::read_example(*pointer, const_cast<char*>(line)));
    }

    VW_EXAMPLE VW_EmptyExample(VW_HANDLE handle) {
        vw * pointer = static_cast<vw*>(handle);
        example* new_ec = VW::new_unused_example(*pointer);
        pointer->p->lp.default_label(new_ec->ld);
        return static_cast<VW_EXAMPLE>(new_ec);
    }

    void VW_ExamplePushFeature(VW_EXAMPLE e, unsigned char ns, uint32_t fid, float v) {
        // warning: assumes namespace exists!
        example * ec = static_cast<example*>(e);
        feature f = { v, fid };
        ec->atomics[ns].push_back(f);
        ec->num_features++;
        ec->sum_feat_sq[ns] += v * v;
        ec->total_sum_feat_sq += v * v;
    }

    void VW_FinishExample(VW_HANDLE handle, VW_EXAMPLE e) {
        vw * pointer = static_cast<vw*>(handle);
        VW::finish_example(*pointer, static_cast<example*>(e));
    }

    float VW_Learn(VW_HANDLE handle, VW_EXAMPLE e) {
        vw * pointer = static_cast<vw*>(handle);
        example * ex = static_cast<example*>(e);
        pointer->learn(ex);
        return VW::get_prediction(ex);
    }

    float VW_GetSimpleLabel(VW_EXAMPLE e) { 
        example * ex = static_cast<example*>(e);
        return ((label_data*)ex->ld)->label; 
    }

    float VW_GetSimpleLabelPrediction(VW_EXAMPLE e) { 
        example * ex = static_cast<example*>(e);
        return ((label_data*)ex->ld)->prediction; 
    }

    float VW_Get_Weight(VW_HANDLE handle, size_t index, size_t offset) {
        vw* pointer = static_cast<vw*>(handle);
        return VW::get_weight(*pointer, (uint32_t) index, (uint32_t) offset);
    }

    void VW_Finish(VW_HANDLE handle) {
        vw * pointer = static_cast<vw*>(handle);
        if (pointer->numpasses > 1) {
            adjust_used_index(*pointer);
            pointer->do_reset_source = true;
            VW::start_parser(*pointer,false);
            pointer->l->driver(pointer);
            VW::end_parser(*pointer); 
        } else {
            release_parser_datastructures(*pointer);
        }

        VW::finish(*pointer);
    }

    float VW_GetPrediction(VW_EXAMPLE e) {
        return VW::get_prediction(static_cast<example*>(e));
    }

    float VW_GetImportance(VW_EXAMPLE e) {
        return VW::get_importance(static_cast<example*>(e));
    }

    size_t VW_HashFeatureA(VW_HANDLE handle, const char * s, unsigned long u) {
        vw * pointer = static_cast<vw*>(handle);
        string str(s);
        return VW::hash_feature(*pointer, str, u);
    }

    size_t VW_GetNumFeatures(VW_EXAMPLE e) { 
        example * ex = static_cast<example*>(e);
        return ex->num_features; 
    }

    float VW_GetLoss(VW_EXAMPLE e) {
        example * ex = static_cast<example*>(e);
        return ex->loss;
    }

    float VW_GetPartialPrediction(VW_EXAMPLE e) { 
        example * ex = static_cast<example*>(e);
        return ex->partial_prediction; 
    }

    /*
    float VW_GetUpdatedPrediction(VW_EXAMPLE e) {
        example * ex = static_cast<example*>(e);
        return ex->updated_prediction;
    }
    */

    /*
    float VW_GetTopicPrediction(VW_EXAMPLE e, size_t i)
    {
        return VW::get_topic_prediction(static_cast<example*>(e), i);
    }
    */

    double VW_GetSumLoss(VW_HANDLE handle) {
        vw * pointer = static_cast<vw*>(handle);
        return pointer->sd->sum_loss;
    }

}
