/* tslint:disable */
/* eslint-disable */
/**
 * SUT SA Example API
 * This is a sample server for SUT SE 2563
 *
 * The version of the OpenAPI document: 1.0
 * Contact: support@swagger.io
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
/**
 * 
 * @export
 * @interface ControllersComplainttype
 */
export interface ControllersComplainttype {
    /**
     * 
     * @type {string}
     * @memberof ControllersComplainttype
     */
    description?: string;
}

export function ControllersComplainttypeFromJSON(json: any): ControllersComplainttype {
    return ControllersComplainttypeFromJSONTyped(json, false);
}

export function ControllersComplainttypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): ControllersComplainttype {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'description': !exists(json, 'description') ? undefined : json['description'],
    };
}

export function ControllersComplainttypeToJSON(value?: ControllersComplainttype | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'description': value.description,
    };
}


