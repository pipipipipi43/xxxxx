// Code generated by protoc-gen-go-http. DO NOT EDIT.
// Source: trace.proto

package pb

import (
	context "context"
	http1 "net/http"
	strings "strings"

	transport "github.com/erda-project/erda-infra/pkg/transport"
	http "github.com/erda-project/erda-infra/pkg/transport/http"
	httprule "github.com/erda-project/erda-infra/pkg/transport/http/httprule"
	runtime "github.com/erda-project/erda-infra/pkg/transport/http/runtime"
	urlenc "github.com/erda-project/erda-infra/pkg/urlenc"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the "github.com/erda-project/erda-infra/pkg/transport/http" package it is being compiled against.
const _ = http.SupportPackageIsVersion1

// TraceServiceHandler is the server API for TraceService service.
type TraceServiceHandler interface {
	// GET /api/msp/apm/traces/{traceID}/spans
	GetSpans(context.Context, *GetSpansRequest) (*GetSpansResponse, error)
	// GET /api/msp/apm/traces
	GetTraces(context.Context, *GetTracesRequest) (*GetTracesResponse, error)
	// GET /api/msp/apm/trace/conditions
	GetTraceQueryConditions(context.Context, *GetTraceQueryConditionsRequest) (*GetTraceQueryConditionsResponse, error)
	// GET /api/msp/apm/trace/debug/histories
	GetTraceDebugHistories(context.Context, *GetTraceDebugHistoriesRequest) (*GetTraceDebugHistoriesResponse, error)
	// GET /api/msp/apm/trace/debug/{requestID}
	GetTraceDebugByRequestID(context.Context, *GetTraceDebugRequest) (*GetTraceDebugResponse, error)
	// POST /api/msp/apm/trace/debug
	CreateTraceDebug(context.Context, *CreateTraceDebugRequest) (*CreateTraceDebugResponse, error)
	// PUT /api/msp/apm/trace/debug/{requestID}
	StopTraceDebug(context.Context, *StopTraceDebugRequest) (*StopTraceDebugResponse, error)
	// GET /api/msp/apm/trace/debug/{requestID}/history/status
	GetTraceDebugHistoryStatusByRequestID(context.Context, *GetTraceDebugStatusByRequestIDRequest) (*GetTraceDebugStatusByRequestIDResponse, error)
}

// RegisterTraceServiceHandler register TraceServiceHandler to http.Router.
func RegisterTraceServiceHandler(r http.Router, srv TraceServiceHandler, opts ...http.HandleOption) {
	h := http.DefaultHandleOptions()
	for _, op := range opts {
		op(h)
	}
	encodeFunc := func(fn func(http1.ResponseWriter, *http1.Request) (interface{}, error)) http.HandlerFunc {
		return func(w http1.ResponseWriter, r *http1.Request) {
			out, err := fn(w, r)
			if err != nil {
				h.Error(w, r, err)
				return
			}
			if err := h.Encode(w, r, out); err != nil {
				h.Error(w, r, err)
			}
		}
	}

	add_GetSpans := func(method, path string, fn func(context.Context, *GetSpansRequest) (*GetSpansResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*GetSpansRequest))
		}
		var GetSpans_info transport.ServiceInfo
		if h.Interceptor != nil {
			GetSpans_info = transport.NewServiceInfo("erda.msp.apm.trace.TraceService", "GetSpans", srv)
			handler = h.Interceptor(handler)
		}
		compiler, _ := httprule.Parse(path)
		temp := compiler.Compile()
		pattern, _ := runtime.NewPattern(httprule.SupportPackageIsVersion1, temp.OpCodes, temp.Pool, temp.Verb)
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, GetSpans_info)
				}
				r = r.WithContext(ctx)
				var in GetSpansRequest
				if err := h.Decode(r, &in); err != nil {
					return nil, err
				}
				var input interface{} = &in
				if u, ok := (input).(urlenc.URLValuesUnmarshaler); ok {
					if err := u.UnmarshalURLValues("", r.URL.Query()); err != nil {
						return nil, err
					}
				}
				params := r.URL.Query()
				if vals := params["scopeId"]; len(vals) > 0 {
					in.ScopeID = vals[0]
				}
				path := r.URL.Path
				if len(path) > 0 {
					components := strings.Split(path[1:], "/")
					last := len(components) - 1
					var verb string
					if idx := strings.LastIndex(components[last], ":"); idx >= 0 {
						c := components[last]
						components[last], verb = c[:idx], c[idx+1:]
					}
					vars, err := pattern.Match(components, verb)
					if err != nil {
						return nil, err
					}
					for k, val := range vars {
						switch k {
						case "traceID":
							in.TraceID = val
						}
					}
				}
				out, err := handler(ctx, &in)
				if err != nil {
					return out, err
				}
				return out, nil
			}),
		)
	}

	add_GetTraces := func(method, path string, fn func(context.Context, *GetTracesRequest) (*GetTracesResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*GetTracesRequest))
		}
		var GetTraces_info transport.ServiceInfo
		if h.Interceptor != nil {
			GetTraces_info = transport.NewServiceInfo("erda.msp.apm.trace.TraceService", "GetTraces", srv)
			handler = h.Interceptor(handler)
		}
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, GetTraces_info)
				}
				r = r.WithContext(ctx)
				var in GetTracesRequest
				if err := h.Decode(r, &in); err != nil {
					return nil, err
				}
				var input interface{} = &in
				if u, ok := (input).(urlenc.URLValuesUnmarshaler); ok {
					if err := u.UnmarshalURLValues("", r.URL.Query()); err != nil {
						return nil, err
					}
				}
				params := r.URL.Query()
				if vals := params["tenantId"]; len(vals) > 0 {
					in.TenantID = vals[0]
				}
				if vals := params["traceId"]; len(vals) > 0 {
					in.TraceID = vals[0]
				}
				out, err := handler(ctx, &in)
				if err != nil {
					return out, err
				}
				return out, nil
			}),
		)
	}

	add_GetTraceQueryConditions := func(method, path string, fn func(context.Context, *GetTraceQueryConditionsRequest) (*GetTraceQueryConditionsResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*GetTraceQueryConditionsRequest))
		}
		var GetTraceQueryConditions_info transport.ServiceInfo
		if h.Interceptor != nil {
			GetTraceQueryConditions_info = transport.NewServiceInfo("erda.msp.apm.trace.TraceService", "GetTraceQueryConditions", srv)
			handler = h.Interceptor(handler)
		}
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, GetTraceQueryConditions_info)
				}
				r = r.WithContext(ctx)
				var in GetTraceQueryConditionsRequest
				if err := h.Decode(r, &in); err != nil {
					return nil, err
				}
				var input interface{} = &in
				if u, ok := (input).(urlenc.URLValuesUnmarshaler); ok {
					if err := u.UnmarshalURLValues("", r.URL.Query()); err != nil {
						return nil, err
					}
				}
				params := r.URL.Query()
				if vals := params["tenantId"]; len(vals) > 0 {
					in.TenantID = vals[0]
				}
				out, err := handler(ctx, &in)
				if err != nil {
					return out, err
				}
				return out, nil
			}),
		)
	}

	add_GetTraceDebugHistories := func(method, path string, fn func(context.Context, *GetTraceDebugHistoriesRequest) (*GetTraceDebugHistoriesResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*GetTraceDebugHistoriesRequest))
		}
		var GetTraceDebugHistories_info transport.ServiceInfo
		if h.Interceptor != nil {
			GetTraceDebugHistories_info = transport.NewServiceInfo("erda.msp.apm.trace.TraceService", "GetTraceDebugHistories", srv)
			handler = h.Interceptor(handler)
		}
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, GetTraceDebugHistories_info)
				}
				r = r.WithContext(ctx)
				var in GetTraceDebugHistoriesRequest
				if err := h.Decode(r, &in); err != nil {
					return nil, err
				}
				var input interface{} = &in
				if u, ok := (input).(urlenc.URLValuesUnmarshaler); ok {
					if err := u.UnmarshalURLValues("", r.URL.Query()); err != nil {
						return nil, err
					}
				}
				params := r.URL.Query()
				if vals := params["terminusKey"]; len(vals) > 0 {
					in.ScopeID = vals[0]
				}
				out, err := handler(ctx, &in)
				if err != nil {
					return out, err
				}
				return out, nil
			}),
		)
	}

	add_GetTraceDebugByRequestID := func(method, path string, fn func(context.Context, *GetTraceDebugRequest) (*GetTraceDebugResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*GetTraceDebugRequest))
		}
		var GetTraceDebugByRequestID_info transport.ServiceInfo
		if h.Interceptor != nil {
			GetTraceDebugByRequestID_info = transport.NewServiceInfo("erda.msp.apm.trace.TraceService", "GetTraceDebugByRequestID", srv)
			handler = h.Interceptor(handler)
		}
		compiler, _ := httprule.Parse(path)
		temp := compiler.Compile()
		pattern, _ := runtime.NewPattern(httprule.SupportPackageIsVersion1, temp.OpCodes, temp.Pool, temp.Verb)
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, GetTraceDebugByRequestID_info)
				}
				r = r.WithContext(ctx)
				var in GetTraceDebugRequest
				if err := h.Decode(r, &in); err != nil {
					return nil, err
				}
				var input interface{} = &in
				if u, ok := (input).(urlenc.URLValuesUnmarshaler); ok {
					if err := u.UnmarshalURLValues("", r.URL.Query()); err != nil {
						return nil, err
					}
				}
				params := r.URL.Query()
				if vals := params["terminusKey"]; len(vals) > 0 {
					in.ScopeID = vals[0]
				}
				path := r.URL.Path
				if len(path) > 0 {
					components := strings.Split(path[1:], "/")
					last := len(components) - 1
					var verb string
					if idx := strings.LastIndex(components[last], ":"); idx >= 0 {
						c := components[last]
						components[last], verb = c[:idx], c[idx+1:]
					}
					vars, err := pattern.Match(components, verb)
					if err != nil {
						return nil, err
					}
					for k, val := range vars {
						switch k {
						case "requestID":
							in.RequestID = val
						}
					}
				}
				out, err := handler(ctx, &in)
				if err != nil {
					return out, err
				}
				return out, nil
			}),
		)
	}

	add_CreateTraceDebug := func(method, path string, fn func(context.Context, *CreateTraceDebugRequest) (*CreateTraceDebugResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*CreateTraceDebugRequest))
		}
		var CreateTraceDebug_info transport.ServiceInfo
		if h.Interceptor != nil {
			CreateTraceDebug_info = transport.NewServiceInfo("erda.msp.apm.trace.TraceService", "CreateTraceDebug", srv)
			handler = h.Interceptor(handler)
		}
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, CreateTraceDebug_info)
				}
				r = r.WithContext(ctx)
				var in CreateTraceDebugRequest
				if err := h.Decode(r, &in); err != nil {
					return nil, err
				}
				var input interface{} = &in
				if u, ok := (input).(urlenc.URLValuesUnmarshaler); ok {
					if err := u.UnmarshalURLValues("", r.URL.Query()); err != nil {
						return nil, err
					}
				}
				out, err := handler(ctx, &in)
				if err != nil {
					return out, err
				}
				return out, nil
			}),
		)
	}

	add_StopTraceDebug := func(method, path string, fn func(context.Context, *StopTraceDebugRequest) (*StopTraceDebugResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*StopTraceDebugRequest))
		}
		var StopTraceDebug_info transport.ServiceInfo
		if h.Interceptor != nil {
			StopTraceDebug_info = transport.NewServiceInfo("erda.msp.apm.trace.TraceService", "StopTraceDebug", srv)
			handler = h.Interceptor(handler)
		}
		compiler, _ := httprule.Parse(path)
		temp := compiler.Compile()
		pattern, _ := runtime.NewPattern(httprule.SupportPackageIsVersion1, temp.OpCodes, temp.Pool, temp.Verb)
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, StopTraceDebug_info)
				}
				r = r.WithContext(ctx)
				var in StopTraceDebugRequest
				if err := h.Decode(r, &in); err != nil {
					return nil, err
				}
				var input interface{} = &in
				if u, ok := (input).(urlenc.URLValuesUnmarshaler); ok {
					if err := u.UnmarshalURLValues("", r.URL.Query()); err != nil {
						return nil, err
					}
				}
				params := r.URL.Query()
				if vals := params["terminusKey"]; len(vals) > 0 {
					in.ScopeID = vals[0]
				}
				path := r.URL.Path
				if len(path) > 0 {
					components := strings.Split(path[1:], "/")
					last := len(components) - 1
					var verb string
					if idx := strings.LastIndex(components[last], ":"); idx >= 0 {
						c := components[last]
						components[last], verb = c[:idx], c[idx+1:]
					}
					vars, err := pattern.Match(components, verb)
					if err != nil {
						return nil, err
					}
					for k, val := range vars {
						switch k {
						case "requestID":
							in.RequestID = val
						}
					}
				}
				out, err := handler(ctx, &in)
				if err != nil {
					return out, err
				}
				return out, nil
			}),
		)
	}

	add_GetTraceDebugHistoryStatusByRequestID := func(method, path string, fn func(context.Context, *GetTraceDebugStatusByRequestIDRequest) (*GetTraceDebugStatusByRequestIDResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*GetTraceDebugStatusByRequestIDRequest))
		}
		var GetTraceDebugHistoryStatusByRequestID_info transport.ServiceInfo
		if h.Interceptor != nil {
			GetTraceDebugHistoryStatusByRequestID_info = transport.NewServiceInfo("erda.msp.apm.trace.TraceService", "GetTraceDebugHistoryStatusByRequestID", srv)
			handler = h.Interceptor(handler)
		}
		compiler, _ := httprule.Parse(path)
		temp := compiler.Compile()
		pattern, _ := runtime.NewPattern(httprule.SupportPackageIsVersion1, temp.OpCodes, temp.Pool, temp.Verb)
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, GetTraceDebugHistoryStatusByRequestID_info)
				}
				r = r.WithContext(ctx)
				var in GetTraceDebugStatusByRequestIDRequest
				if err := h.Decode(r, &in); err != nil {
					return nil, err
				}
				var input interface{} = &in
				if u, ok := (input).(urlenc.URLValuesUnmarshaler); ok {
					if err := u.UnmarshalURLValues("", r.URL.Query()); err != nil {
						return nil, err
					}
				}
				params := r.URL.Query()
				if vals := params["terminusKey"]; len(vals) > 0 {
					in.ScopeID = vals[0]
				}
				path := r.URL.Path
				if len(path) > 0 {
					components := strings.Split(path[1:], "/")
					last := len(components) - 1
					var verb string
					if idx := strings.LastIndex(components[last], ":"); idx >= 0 {
						c := components[last]
						components[last], verb = c[:idx], c[idx+1:]
					}
					vars, err := pattern.Match(components, verb)
					if err != nil {
						return nil, err
					}
					for k, val := range vars {
						switch k {
						case "requestID":
							in.RequestID = val
						}
					}
				}
				out, err := handler(ctx, &in)
				if err != nil {
					return out, err
				}
				return out, nil
			}),
		)
	}

	add_GetSpans("GET", "/api/msp/apm/traces/{traceID}/spans", srv.GetSpans)
	add_GetTraces("GET", "/api/msp/apm/traces", srv.GetTraces)
	add_GetTraceQueryConditions("GET", "/api/msp/apm/trace/conditions", srv.GetTraceQueryConditions)
	add_GetTraceDebugHistories("GET", "/api/msp/apm/trace/debug/histories", srv.GetTraceDebugHistories)
	add_GetTraceDebugByRequestID("GET", "/api/msp/apm/trace/debug/{requestID}", srv.GetTraceDebugByRequestID)
	add_CreateTraceDebug("POST", "/api/msp/apm/trace/debug", srv.CreateTraceDebug)
	add_StopTraceDebug("PUT", "/api/msp/apm/trace/debug/{requestID}", srv.StopTraceDebug)
	add_GetTraceDebugHistoryStatusByRequestID("GET", "/api/msp/apm/trace/debug/{requestID}/history/status", srv.GetTraceDebugHistoryStatusByRequestID)
}
