%global pypi_name apispec

Name:           python-%{pypi_name}
Version:        0.32.0
Release:        1%{?dist}
Summary:        A pluggable API specification generator. Currently supports the OpenAPI specification (f.k.a. Swagger 2.0)

License:        MIT
URL:            https://github.com/marshmallow-code/apispec
Source0:        https://files.pythonhosted.org/packages/source/a/%{pypi_name}/%{pypi_name}-%{version}.tar.gz
BuildArch:      noarch
 
BuildRequires:  python2-devel
BuildRequires:  python-setuptools
 
BuildRequires:  python3-devel
BuildRequires:  python3-setuptools

%description
apispec A pluggable API specification generator. Currently supports the OpenAPI
specification < (f.k.a. Swagger 2.0).Features Supports OpenAPI 2.0
specification (f.k.a. Swagger) Frameworkagnostic Includes plugins for
marshmallow < Flask < Tornado < and bottle < Utilities for parsing
docstringsExample Application .. codeblock:: python from apispec import APISpec
from flask import Flask, jsonify...

%package -n     python3-%{pypi_name}
Summary:        %{summary}
%{?python_provide:%python_provide python3-%{pypi_name}}
 
Requires:       python3-PyYAML >= 3.10
%description -n python3-%{pypi_name}
apispec A pluggable API specification generator. Currently supports the OpenAPI
specification < (f.k.a. Swagger 2.0).Features Supports OpenAPI 2.0
specification (f.k.a. Swagger) Frameworkagnostic Includes plugins for
marshmallow < Flask < Tornado < and bottle < Utilities for parsing
docstringsExample Application .. codeblock:: python from apispec import APISpec
from flask import Flask, jsonify...

%prep
%autosetup -n %{pypi_name}-%{version}
# Remove bundled egg-info
rm -rf %{pypi_name}.egg-info

%build
%py3_build

%install
# Must do the subpackages' install first because the scripts in /usr/bin are
# overwritten with every setup.py install.
%py3_install

%check

%files -n python3-%{pypi_name}
%license docs/license.rst LICENSE
%doc README.rst
%{python3_sitelib}/%{pypi_name}
%{python3_sitelib}/%{pypi_name}-%{version}-py?.?.egg-info

%changelog
* Fri Mar 30 2018 Jan Dobes <jdobes@redhat.com> - 0.32.0-1
- Initial package.
