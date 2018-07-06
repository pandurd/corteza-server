package {package}

{load warning.tpl}

import (
	"net/http"
	"github.com/go-chi/chi"
)

var _ = chi.URLParam

{foreach $calls as $call}
// {name} {call.name} request parameters
type {name|lcfirst}{call.name|capitalize}Request struct {
{foreach $call.parameters as $params}
{foreach $params as $method => $param}
	{param.name} {param.type}{newline}
{/foreach}
{/foreach}
}

func ({name|lcfirst}{call.name|capitalize}Request) new() *{name|lcfirst}{call.name|capitalize}Request {
	return &{name|lcfirst}{call.name|capitalize}Request{}
}

func ({self} *{name|lcfirst}{call.name|capitalize}Request) Fill(r *http.Request) error {
	get := map[string]string{}
	post := map[string]string{}
	urlQuery := r.URL.Query()
	for name, param := range urlQuery {
		get[name] = string(param[0])
	}
	postVars := r.Form
	for name, param := range postVars {
		post[name] = string(param[0])
	}
{foreach $call.parameters as $method => $params}
{foreach $params as $param}
{if strtolower($method) === "path"}
	{self}.{param.name} = chi.URLParam(r, "{param.name}")
{elseif substr($param.type, 0, 2) !== '[]'}
	{self}.{param.name} = {if $param.type !== "string"}{$parsers[$param.type]}({method|strtolower}["{param.name}"]){else}{method|strtolower}["{param.name}"]{/if}{newline}
{/if}
{/foreach}
{/foreach}
	return nil
}

var _ RequestFiller = {name|lcfirst}{call.name|capitalize}Request{}.new()
{/foreach}
